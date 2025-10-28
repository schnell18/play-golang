// Validate XML against XSD in Go using cgo libxml2 binding
package main

import (
	"fmt"
	"os"
	"unsafe"
)

// -----------------------------
// Approach A: using lestrrat-go/libxml2 (recommended)
// -----------------------------
// go get github.com/lestrrat-go/libxml2
// go get github.com/lestrrat-go/libxml2/xsd

/*
Example usage:
  err := ValidateWithLibxml2Go("example.xml", "schema.xsd")
  if err != nil { // validation failed }
*/

// NOTE: keep this code sketchy — check the library docs for exact APIs and error types.

// -----------------------------
// Approach B: using cgo -> libxml2 C API
// -----------------------------
// Requires libxml2 development headers installed (e.g. libxml2-dev on Debian/Ubuntu)
// Build: go build (requires pkg-config + libxml2 available)

/*
#cgo pkg-config: libxml-2.0
#include <libxml/parser.h>
#include <libxml/tree.h>
#include <libxml/xmlschemas.h>
#include <stdlib.h>
*/
import "C"

// ValidateWithCgo validates the XML file at xmlPath against the XSD file at xsdPath using libxml2 C API.
// Returns nil on success or an error describing the validation failure.
func ValidateWithCgo(xmlPath, xsdPath string) error {
	// Read C strings
	xsdC := C.CString(xsdPath)
	defer C.free(unsafe.Pointer(xsdC))

	xmlC := C.CString(xmlPath)
	defer C.free(unsafe.Pointer(xmlC))

	// Initialize libxml
	C.xmlInitParser()
	defer C.xmlCleanupParser()

	// Parse the schema
	schemaParserCtx := C.xmlSchemaNewParserCtxt(xsdC)
	if schemaParserCtx == nil {
		return fmt.Errorf("failed to create schema parser context")
	}
	defer C.xmlSchemaFreeParserCtxt(schemaParserCtx)

	schema := C.xmlSchemaParse(schemaParserCtx)
	if schema == nil {
		return fmt.Errorf("failed to parse XSD schema")
	}
	defer C.xmlSchemaFree(schema)

	// Create validation context
	validCtxt := C.xmlSchemaNewValidCtxt(schema)
	if validCtxt == nil {
		return fmt.Errorf("failed to create schema validation context")
	}
	defer C.xmlSchemaFreeValidCtxt(validCtxt)

	// Parse the XML document
	doc := C.xmlReadFile(xmlC, nil, 0)
	if doc == nil {
		return fmt.Errorf("failed to parse XML document")
	}
	defer C.xmlFreeDoc(doc)

	// Validate
	ret := C.xmlSchemaValidateDoc(validCtxt, doc)
	if ret == 0 {
		// success
		return nil
	}
	// non-zero: validation failed or error
	return fmt.Errorf(
		"validation failed (xmlSchemaValidateDoc returned %d)",
		int(ret),
	)
}

// -----------------------------
// Small helper to show how the libxml2 Go wrapper approach might look
// (This is a sketch — add the real import path and error handling as needed.)
// -----------------------------

/*
func ValidateWithLibxml2Go(xmlPath, xsdPath string) error {
	// Example (pseudocode):
	// import libxml2 "github.com/lestrrat-go/libxml2"
	// import xsd "github.com/lestrrat-go/libxml2/xsd"
	//
	// buf, _ := ioutil.ReadFile(xmlPath)
	// doc, err := libxml2.ParseString(string(buf))
	// defer doc.Free()
	// schema, err := xsd.ParseFromFile(xsdPath)
	// defer schema.Free()
	// if err := schema.Validate(doc); err != nil {
	//     return err
	// }
	// return nil
}
*/

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <xml-file> <xsd-file>\n", os.Args[0])
		os.Exit(2)
	}

	xmlPath := os.Args[1]
	xsdPath := os.Args[2]

	// Example: try CGO approach
	err := ValidateWithCgo(xmlPath, xsdPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Validation failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Validation succeeded")
}
