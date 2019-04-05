NAME = webassembly-tests-go
BUILDDIR = pkg
GOROOT = $(shell go env GOROOT)

default:
	GOOS=js GOARCH=wasm go build -o "${BUILDDIR}/${NAME}.wasm"
	cp "${GOROOT}/misc/wasm/wasm_exec.js" "${BUILDDIR}"

clean:
	rm -f ${NAME} *.wasm
