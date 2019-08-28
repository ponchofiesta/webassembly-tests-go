.PHONY: setup build doc fmt run test clean destroy
# lint vendor_clean vendor_get vendor_update vet

PROJECT=webassembly-benchmarks-go
#export GOPATH=${PWD}/:$(shell go env GOPATH)
export GOROOT=$(shell go env GOROOT)
BUILD_DIR=./build


default: build

build:
	GOOS=js GOARCH=wasm GODEBUG=gcstoptheworld=1 GOGC=20 go build -o "${BUILD_DIR}/${PROJECT}.wasm"
	cp "${GOROOT}/misc/wasm/wasm_exec.js" "${BUILD_DIR}"

doc:
	godoc -http=:6060 -index

fmt:
	go fmt ./...

test:
	go test ./...

clean:
	rm -rf "$BUILD_DIR"
