GOFILES := $(wildcard cmd/*.go) $(wildcard internal/**/*.go)
all:    main

main: $(GOFILES)
	go build cmd/main.go

clean:
	rm -f main
