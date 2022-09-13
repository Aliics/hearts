package dom

import (
	"strings"
)

type ElementAttributes interface {
	Apply(*Element)
}

type Attribute struct {
	Key, Value string
}

func (a Attribute) Apply(e *Element) {
	if strings.Contains(a.Key, ".") {
		var (
			elm  = e
			keys = strings.Split(a.Key, ".")
		)
		for i := 0; i < len(keys)-1; i++ {
			elm = &Element{elm.Get(keys[i])}
		}
		elm.Set(keys[len(keys)-1], a.Value)
	} else {
		e.Set(a.Key, a.Value)
	}
}

type HREFAttribute string

func (h HREFAttribute) Apply(e *Element) {
	e.Set("href", string(h))
}

type PlaceholderAttribute string

func (p PlaceholderAttribute) Apply(e *Element) {
	Attribute{"placeholder", string(p)}.Apply(e)
}

type StyleAttribute string

func (s StyleAttribute) Apply(e *Element) {
	Attribute{"style", string(s)}.Apply(e)
}

type ClassAttribute string

func (c ClassAttribute) Apply(e *Element) {
	Attribute{"class", string(c)}.Apply(e)
}

type MarginAttribute string

func (m MarginAttribute) Apply(e *Element) {
	Attribute{"style.margin", string(m)}.Apply(e)
}

type PaddingAttribute string

func (p PaddingAttribute) Apply(e *Element) {
	Attribute{"style.padding", string(p)}.Apply(e)
}

type WidthAttribute string

func (w WidthAttribute) Apply(e *Element) {
	Attribute{"style.width", string(w)}.Apply(e)
}

type HeightAttribute string

func (h HeightAttribute) Apply(e *Element) {
	Attribute{"style.height", string(h)}.Apply(e)
}

type TitleAttribute string

func (t TitleAttribute) Apply(e *Element) {
	Attribute{"title", string(t)}.Apply(e)
}

type TypeAttribute string

const (
	TypeButton   TypeAttribute = "button"
	TypeCheckbox TypeAttribute = "checkbox"
	TypeColor    TypeAttribute = "color"
	TypeDate     TypeAttribute = "date"
	TypeDatetime TypeAttribute = "datetime-local"
	TypeEmail    TypeAttribute = "email"
	TypeFile     TypeAttribute = "file"
	TypeHidden   TypeAttribute = "hidden"
	TypeImage    TypeAttribute = "image"
	TypeMonth    TypeAttribute = "month"
	TypeNumber   TypeAttribute = "number"
	TypePassword TypeAttribute = "password"
	TypeRadio    TypeAttribute = "radio"
	TypeRange    TypeAttribute = "range"
	TypeReset    TypeAttribute = "reset"
	TypeSearch   TypeAttribute = "search"
	TypeSubmit   TypeAttribute = "submit"
	TypeTel      TypeAttribute = "tel"
	TypeText     TypeAttribute = "text"
	TypeTime     TypeAttribute = "time"
	TypeURL      TypeAttribute = "url"
	TypeWeek     TypeAttribute = "week"
)

func (t TypeAttribute) Apply(e *Element) {
	Attribute{"type", string(t)}.Apply(e)
}

type ValueAttribute string

func (v ValueAttribute) Apply(e *Element) {
	Attribute{"value", string(v)}.Apply(e)
}

type DisplayAttribute string

const (
	DisplayBlock            DisplayAttribute = "block"
	DisplayCompact          DisplayAttribute = "compact"
	DisplayFlex             DisplayAttribute = "flex"
	DisplayGrid             DisplayAttribute = "grid"
	DisplayInline           DisplayAttribute = "inline"
	DisplayInlineBlock      DisplayAttribute = "inline-block"
	DisplayInlineFlex       DisplayAttribute = "inline-flex"
	DisplayInlineTable      DisplayAttribute = "inline-table"
	DisplayListItem         DisplayAttribute = "list-item"
	DisplayMarker           DisplayAttribute = "marker"
	DisplayNone             DisplayAttribute = "none"
	DisplayRunIn            DisplayAttribute = "run-in"
	DisplayTable            DisplayAttribute = "table"
	DisplayTableCaption     DisplayAttribute = "table-caption"
	DisplayTableCell        DisplayAttribute = "table-cell"
	DisplayTableColumn      DisplayAttribute = "table-column"
	DisplayTableColumnGroup DisplayAttribute = "table-column-group"
	DisplayTableFooterGroup DisplayAttribute = "table-footer-group"
	DisplayTableHeaderGroup DisplayAttribute = "table-header-group"
	DisplayTableRow         DisplayAttribute = "table-row"
	DisplayTableRowGroup    DisplayAttribute = "table-row-group"
	DisplayInitial          DisplayAttribute = "initial"
	DisplayInherit          DisplayAttribute = "inherit"
)

func (d DisplayAttribute) Apply(e *Element) {
	Attribute{"style.display", string(d)}.Apply(e)
}

type FlexDirectionAttribute string

const (
	FlexDirectionRow           FlexDirectionAttribute = "row"
	FlexDirectionRowReverse    FlexDirectionAttribute = "row-reverse"
	FlexDirectionColumn        FlexDirectionAttribute = "column"
	FlexDirectionColumnReverse FlexDirectionAttribute = "column-reverse"
	FlexDirectionInitial       FlexDirectionAttribute = "initial"
	FlexDirectionInherit       FlexDirectionAttribute = "inherit"
)

func (f FlexDirectionAttribute) Apply(e *Element) {
	Attribute{"style.flex-direction", string(f)}.Apply(e)
}
