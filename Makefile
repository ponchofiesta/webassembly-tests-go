.PHONY: setup build doc fmt run test clean destroy
# lint vendor_clean vendor_get vendor_update vet

GO ?= go1.12.9
PROJECT=webassembly-benchmarks-go
#export GOPATH=${PWD}/:$(shell go env GOPATH)
export GOROOT=$(shell $(GO) env GOROOT)
BUILD_DIR=./build


default: build

build:
	GOOS=js GOARCH=wasm GODEBUG=gcstoptheworld=1 GOGC=20 ${GO} build -o "${BUILD_DIR}/${PROJECT}.wasm"
	cp "${GOROOT}/misc/wasm/wasm_exec.js" "${BUILD_DIR}"

doc:
	godoc -http=:6060 -index

fmt:
	${GO} fmt ./...

test:
	${GO} test ./...

clean:
	rm -rf "$BUILD_DIR"
