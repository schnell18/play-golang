# Introduction

Explores [the flight recorder][1] in go 1.25.

## Experiment steps

Run the main program:

    go run main.go

In a separate Window, create loads for the API:

    sh test.sh

Identify the `snapshot.trace` file and run analysis
and virtualisation:

    go tool trace snapshot.trace


[1]: https://go.dev/blog/flight-recorder
