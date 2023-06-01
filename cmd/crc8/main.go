package main

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/sigurn/crc8"
)

func main() {
	ver1 := "v1.2.31-beta1+b01"
	fmt.Printf("%s as big int => %d\n", ver1, toBigInt(ver1))
	ver1 = "v1.2.31"
	fmt.Printf("%s as big int => %d\n", ver1, toBigInt(ver1))
	ver1 = "v2.2.31-beta1"
	fmt.Printf("%s as big int => %d\n", ver1, toBigInt(ver1))
	ver1 = "v2.2.31-beta1+b01"
	fmt.Printf("%s as big int => %d\n", ver1, toBigInt(ver1))
	ver1 = "v2.2.31-beta1+sha12342424244"
	fmt.Printf("%s as big int => %d\n", ver1, toBigInt(ver1))
	ver1 = "v1.20.31-RC01"
	fmt.Printf("%s as big int => %d\n", ver1, toBigInt(ver1))
	ver1 = "1.20.31-RC01"
	fmt.Printf("%s as big int => %d\n", ver1, toBigInt(ver1))
	ver1 = "1.x.31-RC01"
	fmt.Printf("%s as big int => %d\n", ver1, toBigInt(ver1))
}

func toBigInt(vstr string) uint64 {
	var other string
	v, err := semver.NewVersion(vstr)
	if err != nil {
		return 0
	}
	other = v.Prerelease()
	var crc uint8
	if other != "" {
		table := crc8.MakeTable(crc8.CRC8_MAXIM)
		crc = crc8.Checksum([]byte(other), table)
	}

	return v.Major()*100*100*10000 + v.Minor()*100*10000 + v.Patch()*10000 + uint64(crc)
}
