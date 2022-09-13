package dom

import "syscall/js"

var (
	Document     = valueAsElement(js.Global().Get("document"))
	DocumentBody = valueAsElement(Document.Get("body"))
)

func NewWebSocket(url string) Element {
	return Element{js.Global().Get("WebSocket").New(url)}
}

func NewXHR(method, url string) Element {
	req := Element{js.Global().Get("XMLHttpRequest").New()}
	req.Call("open", method, url)
	return req
}

func valueAsElement(value js.Value) Element {
	return Element{value}
}
