#!/bin/bash

for j in $(seq 1 100); do
    for i in $(seq 1 100); do
        curl http://localhost:8090/guess-number\?guess=$i
    done
done
