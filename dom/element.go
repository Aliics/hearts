package dom

import (
	"syscall/js"
)

type Element struct {
	js.Value
}

func CreateElement(elementType ElementType) Element {
	return Element{Document.Call("createElement", string(elementType))}
}

func (e Element) AppendChild(child Element) Element {
	e.Call("appendChild", child.Value)
	return e
}

func (e Element) AppendChildren(children ...Element) Element {
	for _, child := range children {
		e.AppendChild(child)
	}
	return e
}

func (e Element) AddEventListener(eventType EventType, f func(Element)) Element {
	eventListenerFunc := js.FuncOf(func(value js.Value, _ []js.Value) any {
		f(e)
		return js.Undefined()
	})
	e.Call("addEventListener", string(eventType), eventListenerFunc)
	return e
}
