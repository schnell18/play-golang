#!/bin/bash

go build -ldflags="-X 'main.Version=v1.0.0' -X 'github.com/schnell18/play-golang/ldflags-ver/build.User=Justin'"
# go build -ldflags="-X 'main.Version=v1.0.0' -X 'build.User=Justin'"
