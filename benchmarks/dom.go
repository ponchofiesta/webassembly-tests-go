package benchmarks

import (
	"strconv"
	"syscall/js"
)

func Dom(n int) js.Value {
	window := js.Global()
	document := window.Get("document")
	body := document.Get("body")
	container := document.Call("createElement", "div")
	body.Call("appendChild", container)

	for i := 0; i < 10000; i++ {
		header := document.Call("createElement", "h3")
		header.Set("textContent", "Header "+strconv.Itoa(i))
		paragraph := document.Call("createElement", "p")
		paragraph.Set("textContent", "Paragraph "+strconv.Itoa(i))
		container.Call("appendChild", header)
		container.Call("appendChild", paragraph)
	}

	body.Call("removeChild", container)

	return js.Undefined()
}
