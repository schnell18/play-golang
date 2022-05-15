# Introduction

Explore [brotli][1] compression.
Run an instance of nginx w/ brotli enabled.
Write a golang program make request to nginx with `br` encoding.

Compare brotli and lz4.
lz4 seems unable to compress small data.
brotli compress level 9 get 3:1 compress ratio.


## setup

To run nginx w/ brotli enabled, you launch container like:

    docker run -p 8000:80 fholzer/nginx-brotli:v1.21.6

[1]: https://github.com/google/brotli
