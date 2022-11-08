# celo-tx-decoder

## Get started

1. Bootstrap the project
```bash
make
```

2. Build the docker image
```bash
make image
```

## Issues

The image does not build because of build constraints on the package `github.com/celo-org/celo-bls-go`.
```
Step 7/14 : RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -installsuffix nocgo -o /app cmd/main.go
 ---> Running in 88a8d8b27a6f
package command-line-arguments
        imports github.com/celo-org/celo-blockchain/core/types
        imports github.com/celo-org/celo-blockchain/crypto/bls
        imports github.com/celo-org/celo-bls-go/bls: build constraints exclude all Go files in /proxy/vendor/github.com/celo-org/celo-bls-go/bls
```
It seems to be related to `C` dependencies. More details [here](https://github.com/celo-org/celo-bls-go-linux/blob/49606a0c333d28f73f77c3c02634158fa595112a/bls/bls.go#L6)
