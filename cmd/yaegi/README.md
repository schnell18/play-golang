# Run yaegi as REPL

Install yaegi cli:

    go install github.com/traefik/yaegi/cmd/yaegi@latest

# Run go as shell script

~~~shell
#!/usr/bin/env yaegi
package main

import "fmt"

func main() {
    fmt.Println("I'm a Go script!")
}
~~~

# Alternatives

- Go Plugin
- Lua
- WASM
- gopher-lua
