package main

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/google/brotli/go/cbrotli"
	"github.com/pierrec/lz4"
)

func cbrotliEncode(toCompress []byte) ([]byte, error) {
	opts := cbrotli.WriterOptions{
		Quality: 9,
		LGWin:   0,
	}
	encoded, err := cbrotli.Encode(toCompress, opts)
	if err != nil {
		log.Fatalf("brotli encode failed due to: %s", err)
	}
	fmt.Printf("Brotli: %v\n", encoded)
	fmt.Printf("Brotli: after compression bytes: %d\n", len(encoded))
	fmt.Printf("Brotli: %0x\n", encoded)
	return encoded, err
}

func cbrotliDecode(toDecompress []byte) ([]byte, error) {
	decoded, err := cbrotli.Decode(toDecompress)
	if err != nil {
		log.Fatalf("brotli decode failed due to: %s", err)
	}
	fmt.Printf("Brotli: %s\n", string(decoded))
	return decoded, err
}

func lz4Encode(toCompress []byte) ([]byte, error) {

	buf := make([]byte, len(toCompress))
	ht := make([]int, 64<<10) // buffer for the compression table

	n, err := lz4.CompressBlock(toCompress, buf, ht)
	if err != nil {
		log.Fatalf("lz4 decode failed due to: %s", err)
	}
	buf = buf[:n] // compressed data
	fmt.Printf("lz4: %v\n", buf)
	fmt.Printf("lz4: after compression bytes: %d\n", len(buf))
	fmt.Printf("lz4: %0x\n", buf)
	return buf, err
}

func lz4Decode(toDecompress []byte) ([]byte, error) {

	buf := make([]byte, len(toDecompress))

	// Allocated a very large buffer for decompression.
	out := make([]byte, 10*len(toDecompress))
	n, err := lz4.UncompressBlock(buf, out)
	if err != nil {
		log.Fatalf("lz4 decode failed due to: %s", err)
	}
	out = out[:n] // uncompressed data

	fmt.Printf("lz4: %s\n", string(out))
	return out, err
}

func restoreBits(decoded []byte, source string) {
	restoredBits := make([]uint32, 0, 64)
	for i := 0; i < len(decoded); i += 4 {
		sec := decoded[i : i+4]
		restoredBits = append(restoredBits, binary.LittleEndian.Uint32(sec))
	}
	fmt.Printf("%s: %v\n", source, restoredBits)

}

func main() {
	bits := make([]uint32, 64)
	src := make([]byte, 4*len(bits))
	for i := range bits {
		bits[i] = uint32(i * 2)
		buf := src[i*4 : (i+1)*4]
		binary.LittleEndian.PutUint32(buf, bits[i])
	}
	fmt.Printf("before compression bytes: %d\n", len(bits)*4)
	fmt.Printf("%v\n", bits)

	encoded, _ := lz4Encode(src)
	decoded, _ := lz4Decode(encoded)
	restoreBits(decoded, "lz4")

	encoded, _ = cbrotliEncode(src)
	decoded, _ = cbrotliDecode(encoded)
	restoreBits(decoded, "brotli")

}
