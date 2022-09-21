package web

import (
	. "github.com/aliics/hearts/web/dom"
)

func centeredModal(children ...Element) Element {
	return Div(
		DisplayFlex,
		FlexDirectionColumn,
		WidthAttribute("fit-content"),
		PaddingAttribute("1em"),
		PositionAttribute("absolute"),
		StyleAttribute("top: 50%; left: 50%; transform: translate(-50%, -50%);"),
		BorderRadiusAttribute("4px"),
		BackgroundColorAttribute("#f1f1f1"),
	)(children...)
}
