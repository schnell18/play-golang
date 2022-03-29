package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/howeyc/crc16"
)

func main() {
	scanner := bufio.NewScanner((os.Stdin))
	for scanner.Scan() {
		txt := scanner.Text()
		checksum := crc16.ChecksumCCITT([]byte(txt))
		fmt.Printf("%s: %d\n", txt, checksum)

	}
}
