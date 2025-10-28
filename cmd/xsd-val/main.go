// xsd_validation.go
// Fixed: working cgo example that compiles and returns structured validation errors

package main

/*
#cgo pkg-config: libxml-2.0
#include <libxml/parser.h>
#include <libxml/tree.h>
#include <libxml/xmlschemas.h>
#include <stdlib.h>
#include <stdarg.h>
#include <string.h>

// Error callback: append formatted message into ctx (a char buffer)
static void schemaErrorCallback(void *ctx, const char *msg, ...) {
    if (ctx == NULL) return;
    char *buf = (char*)ctx;
    size_t cur = strlen(buf);
    if (cur >= 8191) return; // no space

    va_list ap;
    va_start(ap, msg);
    vsnprintf(buf + cur, 8192 - cur, msg, ap);
    va_end(ap);
}

// Helper to set both error and warning callbacks (avoids passing function pointers from Go)
static void set_schema_errors(xmlSchemaValidCtxtPtr valid_ctxt, void *ctx) {
    xmlSchemaSetValidErrors(valid_ctxt, (xmlSchemaValidityErrorFunc)schemaErrorCallback, (xmlSchemaValidityWarningFunc)schemaErrorCallback, ctx);
}
*/
import "C"

import (
	"fmt"
	"os"
	"strings"
	"unsafe"
)

// ValidationError aggregates validation messages
type ValidationError struct {
	Errors []string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf(
		"%d validation error(s):\n%s",
		len(e.Errors),
		strings.Join(e.Errors, "\n"),
	)
}

// ValidateWithCgo validates an XML file against an XSD file using libxml2 C API.
// Returns nil on success or *ValidationError on validation problems.
func ValidateWithCgo(xmlPath, xsdPath string) error {
	xsdC := C.CString(xsdPath)
	defer C.free(unsafe.Pointer(xsdC))

	xmlC := C.CString(xmlPath)
	defer C.free(unsafe.Pointer(xmlC))

	C.xmlInitParser()
	defer C.xmlCleanupParser()

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

	validCtxt := C.xmlSchemaNewValidCtxt(schema)
	if validCtxt == nil {
		return fmt.Errorf("failed to create schema validation context")
	}
	defer C.xmlSchemaFreeValidCtxt(validCtxt)

	// Prepare buffer for C callback to write into
	buf := make([]byte, 8192) // zero-initialized

	// set callbacks via helper function defined in C above
	C.set_schema_errors(validCtxt, unsafe.Pointer(&buf[0]))

	doc := C.xmlReadFile(xmlC, nil, 0)
	if doc == nil {
		return fmt.Errorf("failed to parse XML document")
	}
	defer C.xmlFreeDoc(doc)

	ret := C.xmlSchemaValidateDoc(validCtxt, doc)

	// Extract string up to first NUL
	n := 0
	for ; n < len(buf); n++ {
		if buf[n] == 0 {
			break
		}
	}
	msg := strings.TrimSpace(string(buf[:n]))

	if ret != 0 || msg != "" {
		var lines []string
		if msg != "" {
			for l := range strings.SplitSeq(msg, "\n") {
				l = strings.TrimSpace(l)
				if l != "" {
					lines = append(lines, l)
				}
			}
		}
		// If xmlSchemaValidateDoc returned non-zero but no message was captured, add a generic message
		if len(lines) == 0 {
			lines = append(
				lines,
				fmt.Sprintf(
					"xmlSchemaValidateDoc returned %d (validation failed)",
					int(ret),
				),
			)
		}
		return &ValidationError{Errors: lines}
	}

	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <xml-file> <xsd-file>\n", os.Args[0])
		os.Exit(2)
	}

	xmlPath := os.Args[1]
	xsdPath := os.Args[2]

	err := ValidateWithCgo(xmlPath, xsdPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Validation failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Validation succeeded")
}
