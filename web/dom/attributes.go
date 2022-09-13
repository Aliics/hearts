package dom

type ElementAttributes interface {
	Apply(*Element)
}

type HREF string

func (h HREF) Apply(e *Element) {
	e.Set("href", string(h))
}

type Placeholder string

func (p Placeholder) Apply(e *Element) {
	e.Set("placeholder", string(p))
}

type Style string

func (s Style) Apply(e *Element) {
	e.Set("style", string(s))
}

type Class string

func (c Class) Apply(e *Element) {
	e.Set("style", string(c))
}

type Margin string

func (m Margin) Apply(e *Element) {
	e.Get("style").Set("margin", string(m))
}

type Padding string

func (p Padding) Apply(e *Element) {
	e.Get("style").Set("padding", string(p))
}

type Width string

func (w Width) Apply(e *Element) {
	e.Get("style").Set("width", string(w))
}

type Height string

func (h Height) Apply(e *Element) {
	e.Get("style").Set("height", string(h))
}

type Type string

const (
	TypeButton   Type = "button"
	TypeCheckbox Type = "checkbox"
	TypeColor    Type = "color"
	TypeDate     Type = "date"
	TypeDatetime Type = "datetime-local"
	TypeEmail    Type = "email"
	TypeFile     Type = "file"
	TypeHidden   Type = "hidden"
	TypeImage    Type = "image"
	TypeMonth    Type = "month"
	TypeNumber   Type = "number"
	TypePassword Type = "password"
	TypeRadio    Type = "radio"
	TypeRange    Type = "range"
	TypeReset    Type = "reset"
	TypeSearch   Type = "search"
	TypeSubmit   Type = "submit"
	TypeTel      Type = "tel"
	TypeText     Type = "text"
	TypeTime     Type = "time"
	TypeURL      Type = "url"
	TypeWeek     Type = "week"
)

func (t Type) Apply(e *Element) {
	e.Set("type", string(t))
}

type Value string

func (v Value) Apply(e *Element) {
	e.Set("value", string(v))
}

type Display string

const (
	DisplayBlock            Display = "block"
	DisplayCompact          Display = "compact"
	DisplayFlex             Display = "flex"
	DisplayGrid             Display = "grid"
	DisplayInline           Display = "inline"
	DisplayInlineBlock      Display = "inline-block"
	DisplayInlineFlex       Display = "inline-flex"
	DisplayInlineTable      Display = "inline-table"
	DisplayListItem         Display = "list-item"
	DisplayMarker           Display = "marker"
	DisplayNone             Display = "none"
	DisplayRunIn            Display = "run-in"
	DisplayTable            Display = "table"
	DisplayTableCaption     Display = "table-caption"
	DisplayTableCell        Display = "table-cell"
	DisplayTableColumn      Display = "table-column"
	DisplayTableColumnGroup Display = "table-column-group"
	DisplayTableFooterGroup Display = "table-footer-group"
	DisplayTableHeaderGroup Display = "table-header-group"
	DisplayTableRow         Display = "table-row"
	DisplayTableRowGroup    Display = "table-row-group"
	DisplayInitial          Display = "initial"
	DisplayInherit          Display = "inherit"
)

func (d Display) Apply(e *Element) {
	e.Get("style").Set("display", string(d))
}

type FlexDirection string

const (
	FlexDirectionRow           FlexDirection = "row"
	FlexDirectionRowReverse    FlexDirection = "row-reverse"
	FlexDirectionColumn        FlexDirection = "column"
	FlexDirectionColumnReverse FlexDirection = "column-reverse"
	FlexDirectionInitial       FlexDirection = "initial"
	FlexDirectionInherit       FlexDirection = "inherit"
)

func (f FlexDirection) Apply(e *Element) {
	e.Get("style").Set("flex-direction", string(f))
}
