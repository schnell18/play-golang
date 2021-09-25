package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var alg = flag.String("a", "", "checksum type, sha256/sha384/sha512 etc")

func main() {
	flag.Parse()
	if *alg == "" {
		*alg = "sha256"
	}

	if *alg != "sha256" && *alg != "sha384" && *alg != "sha512" {
		fmt.Fprintf(os.Stderr, "bad algorithm: %s, valid algorithm are sha256, sha384 or sha512\n", *alg)
		os.Exit(1)
	}

	buf := make([]byte, 2)
	n, err := os.Stdin.Read(buf)
	if err == nil {
		switch *alg {
		case "sha256":
			fmt.Printf("%x\n", sha256.Sum256(buf[:n]))
		case "sha384":
			fmt.Printf("%x\n", sha512.Sum384(buf[:n]))
		case "sha512":
			fmt.Printf("%x\n", sha512.Sum512(buf[:n]))
		}
	}

}
