# Amazon S3 client-side encryption demo

This is a simple demo of managing AES key with [OpenBao][1].


## Setup a local OpenBao dev server

Install OpenBao using package manager. For example:

    yay openbao

Launch an OpenBao dev server:

    bao server -dev -dev-root-token-id="localenv"

Provision a AES secret:

    export BAO_ADDR=127.0.0.1:8200
    export BAO_TOKEN=localenv
    bao kv put secret/my-aes-key key="v7sRMFIhb0RAtKvik08L6XycbMgWK3Y4XIU5iB4ytLs="
    bao kv get --format json secret/my-aes-key | jq -r ".data.data.key"

## Run the main.go

    go run ./main.go

[1]: https://openbao.org/
