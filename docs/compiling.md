# Compiling

Compiling from scratch requires the [Go programming language toolchain](https://golang.org/dl/) and git. Note: *applehealth2csv* uses [go modules](https://github.com/golang/go/wiki/Modules) for dependency management.

To generate native binary, type 

```
go build
```

Please look at `Makefile` for cross-compiling for other platforms.
