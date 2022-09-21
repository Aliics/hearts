package dom

import "syscall/js"

type ElementType string

const (
	ElementTypeA        ElementType = "a"
	ElementTypeABBR     ElementType = "abbr"
	ElementTypeAddress  ElementType = "address"
	ElementTypeArea     ElementType = "area"
	ElementTypeB        ElementType = "b"
	ElementTypeBody     ElementType = "body"
	ElementTypeBR       ElementType = "br"
	ElementTypeButton   ElementType = "button"
	ElementTypeCode     ElementType = "code"
	ElementTypeCol      ElementType = "col"
	ElementTypeColGroup ElementType = "colgroup"
	ElementTypeDiv      ElementType = "div"
	ElementTypeFooter   ElementType = "footer"
	ElementTypeH1       ElementType = "h1"
	ElementTypeH2       ElementType = "h2"
	ElementTypeH3       ElementType = "h3"
	ElementTypeH4       ElementType = "h4"
	ElementTypeH5       ElementType = "h5"
	ElementTypeH6       ElementType = "h6"
	ElementTypeHead     ElementType = "head"
	ElementTypeHeader   ElementType = "header"
	ElementTypeInput    ElementType = "input"
	ElementTypeImage    ElementType = "img"
	ElementTypeIFrame   ElementType = "iframe"
	ElementTypeLI       ElementType = "li"
	ElementTypeP        ElementType = "p"
	ElementTypePre      ElementType = "pre"
	ElementTypeSection  ElementType = "section"
	ElementTypeSpan     ElementType = "span"
	ElementTypeTable    ElementType = "table"
	ElementTypeTD       ElementType = "td"
	ElementTypeTHead    ElementType = "thead"
	ElementTypeTitle    ElementType = "title"
	ElementTypeTR       ElementType = "tr"
	ElementTypeUL       ElementType = "ul"
)

func StringLiteral(s string) Element {
	return Element{js.ValueOf(s)}
}

func A(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeA, attributes)
}

func ABBR(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeABBR, attributes)
}

func Address(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeAddress, attributes)
}

func Area(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeArea, attributes)
}

func B(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeB, attributes)
}

func Body(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeBody, attributes)
}

func BR(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeBR, attributes)
}

func Button(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeButton, attributes)
}

func Code(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeCode, attributes)
}

func Col(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeCol, attributes)
}

func Comment(text string) Element {
	return Element{Document.Call("createComment", text)}
}

func Group(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeColGroup, attributes)
}

func Div(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeDiv, attributes)
}

func Empty() Element {
	return Comment("empty")
}

func Footer(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeFooter, attributes)
}
func H1(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeH1, attributes)
}
func H2(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeH2, attributes)
}
func H3(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeH3, attributes)
}
func H4(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeH4, attributes)
}
func H5(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeH5, attributes)
}
func H6(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeH6, attributes)
}
func Head(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeHead, attributes)
}
func Header(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeHeader, attributes)
}

func Input(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeInput, attributes)
}

func Image(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeImage, attributes)
}

func IFrame(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeIFrame, attributes)
}

func LI(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeLI, attributes)
}

func P(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeP, attributes)
}

func Pre(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypePre, attributes)
}

func Section(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeSection, attributes)
}

func Span(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeSpan, attributes)
}

func Table(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeTable, attributes)
}

func TD(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeTD, attributes)
}

func THead(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeTHead, attributes)
}

func Title(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeTitle, attributes)
}

func TR(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeTR, attributes)
}

func UL(attributes ...ElementAttribute) func(...Element) Element {
	return attributeElement(ElementTypeUL, attributes)
}

func attributeElement(elementType ElementType, attributes []ElementAttribute) func(...Element) Element {
	return func(inner ...Element) Element {
		e := CreateElement(elementType)
		for _, attribute := range attributes {
			if attribute == nil {
				continue
			}
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
