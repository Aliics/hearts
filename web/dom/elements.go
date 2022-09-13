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

func A(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeA, attributes)
}

func ABBR(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeABBR, attributes)
}

func Address(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeAddress, attributes)
}

func Area(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeArea, attributes)
}

func B(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeB, attributes)
}

func Body(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeBody, attributes)
}

func BR(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeBR, attributes)
}

func Button(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeButton, attributes)
}

func Code(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeCode, attributes)
}

func Col(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeCol, attributes)
}

func Group(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeColGroup, attributes)
}

func Div(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeDiv, attributes)
}

func Footer(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeFooter, attributes)
}
func H1(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeH1, attributes)
}
func H2(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeH2, attributes)
}
func H3(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeH3, attributes)
}
func H4(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeH4, attributes)
}
func H5(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeH5, attributes)
}
func H6(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeH6, attributes)
}
func Head(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeHead, attributes)
}
func Header(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeHeader, attributes)
}

func Input(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeInput, attributes)
}

func Image(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeImage, attributes)
}

func IFrame(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeIFrame, attributes)
}

func LI(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeLI, attributes)
}

func P(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeP, attributes)
}

func Pre(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypePre, attributes)
}

func Section(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeSection, attributes)
}

func Span(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeSpan, attributes)
}

func Table(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeTable, attributes)
}

func TD(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeTD, attributes)
}

func THead(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeTHead, attributes)
}

func Title(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeTitle, attributes)
}

func TR(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeTR, attributes)
}

func UL(attributes ...ElementAttributes) func(...Element) Element {
	return attributeElement(ElementTypeUL, attributes)
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
