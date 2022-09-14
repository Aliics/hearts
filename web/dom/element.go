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

func (e Element) RemoveChild(child Element) Element {
	e.Call("removeChild", child.Value)
	return e
}

func (e Element) RemoveChildren(children ...Element) Element {
	for _, child := range children {
		e.RemoveChild(child)
	}
	return e
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

func (e Element) AddEventListener(eventType EventType, f func(js.Value, []js.Value)) Element {
	eventListenerFunc := js.FuncOf(func(value js.Value, args []js.Value) any {
		f(value, args)
		return js.Undefined()
	})
	e.Call("addEventListener", string(eventType), eventListenerFunc)
	return e
}
