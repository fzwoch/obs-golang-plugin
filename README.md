# OBS Studio Golang Plugin

An [OBS Studio] Plugin template in [Go]. Start writing OBS plugins in Go. For whatever reason you may have.

This library hooks into the C plugin interface of OBS via Go's [cgo] system.

[OBS Studio]: https://obsproject.com/
[Go]: https://golang.org/
[cgo]: https://golang.org/cmd/cgo/

## Build

    go build -buildmode=c-shared -o plugin.[so|dll|dylib]
