.PHONY: setup build doc fmt run test clean destroy
# lint vendor_clean vendor_get vendor_update vet

GO ?= go1.12.9
PROJECT=webassembly-benchmarks-go
#export GOPATH=${PWD}/:$(shell go env GOPATH)
export GOROOT=$(shell $(GO) env GOROOT)
export GOOS=js
export GOARCH=wasm
export GOGC=20

ifeq ($(OS),Windows_NT)
	BUILD_DIR=.\build
else
	BUILD_DIR=./build
endif

default: build

build:
ifeq ($(OS),Windows_NT)
    ${GO} build -o "${BUILD_DIR}\${PROJECT}.wasm"
	copy "${GOROOT}\misc\wasm\wasm_exec.js" "${BUILD_DIR}"
else
    ${GO} build -o "${BUILD_DIR}/${PROJECT}.wasm"
	cp "${GOROOT}/misc/wasm/wasm_exec.js" "${BUILD_DIR}"
endif

doc:
	godoc -http=:6060 -index

fmt:
	${GO} fmt ./...

test:
	${GO} test ./...

clean:
	rm -rf "$BUILD_DIR"
