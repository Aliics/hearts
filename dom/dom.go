package dom

import "syscall/js"

var (
	Document = valueAsElement(js.Global().Get("document"))
	Body     = valueAsElement(Document.Get("body"))
)

func valueAsElement(value js.Value) Element {
	return Element{value}
}
