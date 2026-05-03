package main

import (
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func main() {
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)

	_, err := i.Eval(`import "fmt"`)
	if err != nil {
		panic(err)
	}

	_, err = i.Eval(`fmt.Println("Hello Yaegi")`)
	if err != nil {
		panic(err)
	}
}
