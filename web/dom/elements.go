package dom

import "syscall/js"

type ElementType string

const (
	ElementTypeA       ElementType = "a"
	ElementTypeP       ElementType = "p"
	ElementTypeInput   ElementType = "input"
	ElementTypeDiv     ElementType = "div"
	ElementTypeSection ElementType = "section"
	ElementTypeTable   ElementType = "table"
	ElementTypeTHead   ElementType = "thead"
	ElementTypeTR      ElementType = "tr"
	ElementTypeTD      ElementType = "td"
	ElementTypeUL      ElementType = "ul"
	ElementTypeLI      ElementType = "li"
)

func StringLiteral(s string) Element {
	return Element{js.ValueOf(s)}
}

func A(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeA, attributes)
}

func P(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeP, attributes)
}

func Input(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeInput, attributes)
}

func Div(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeDiv, attributes)
}

func Section(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeSection, attributes)
}

func Table(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeTable, attributes)
}

func THead(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeTHead, attributes)
}

func TR(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeTR, attributes)
}

func TD(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeTD, attributes)
}

func UL(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeUL, attributes)
}

func LI(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeLI, attributes)
}

func attributeElement(elementType ElementType, attributes []ElementAttributes) func(...Element) Element {
	return func(inner ...Element) Element {
		e := CreateElement(elementType)
		for _, attribute := range attributes {
			attribute.Apply(&e)
		}
		for _, element := range inner {
			if element.Type() == js.TypeString {
				e.Set("innerHTML", e.Get("innerHTML").String()+element.String())
			} else {
				e.AppendChild(element)
			}
		}
		return e
	}
}
