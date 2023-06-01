package main

import (
	"fmt"
	"sort"

	"github.com/Masterminds/semver/v3"
)

func main() {
	raw := []string{"1.2.3", "1.0", "1.3", "2", "0.4.2", "0.4.2-beta"}
	vs := make([]*semver.Version, len(raw))
	for i, r := range raw {
		v, err := semver.NewVersion(r)
		if err != nil {
			fmt.Printf("Error parsing version: %s", err)
		}

		vs[i] = v
	}
	fmt.Printf("before sort: %v\n", vs)
	sort.Sort(semver.Collection(vs))
	fmt.Printf("after sort: %v\n", vs)
}
