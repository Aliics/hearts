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

func (e Element) ReplacedWithEmpty() Element {
	return e.Replaced(Empty())
}

func (e Element) Replaced(new Element) Element {
	if e.IsNull() || e.IsUndefined() {
		return e
	}

	parent := e.Get("parentNode")
	if parent.IsNull() || parent.IsUndefined() {
		return e
	}

	Element{parent}.ReplaceChild(e, new)
	return new
}

func (e Element) ReplaceChild(old, new Element) Element {
	childNodes := e.Get("childNodes")
	for i := 0; i < childNodes.Length(); i++ {
		if childNodes.Index(i).Equal(old.Value) {
			e.Call("replaceChild", new.Value, old.Value)
			return e
		}
	}

	return e
}

func (e Element) RemoveChild(child Element) Element {
	childNodes := e.Get("childNodes")
	for i := 0; i < childNodes.Length(); i++ {
		if childNodes.Index(i).Equal(child.Value) {
			e.Call("removeChild", child.Value)
			return e
		}
	}

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
