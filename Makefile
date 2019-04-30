.PHONY: setup build doc fmt run test clean destroy
# lint vendor_clean vendor_get vendor_update vet

PROJECT=webassembly-benchmarks-go

export GOPATH=${PWD}/:$(shell go env GOROOT)
export GOROOT=$(shell go env GOROOT)


default: build

build:
	GOOS=js GOARCH=wasm GODEBUG=gcstoptheworld=1 GOGC=20 go build -o ./bin/${PROJECT}.wasm ./src
	cp "${GOROOT}/misc/wasm/wasm_exec.js" ./bin

doc:
	godoc -http=:6060 -index

fmt:
	go fmt ./src/...

test:
	go test ./src/tests/...

clean:
	rm -rf ./bin
