//go:build wasm && js

package web

import (
	. "github.com/aliics/hearts/web/dom"
)

var (
	gameScreen = Empty()
	menuScreen = Empty()
)

func Run() {
	menuScreen = createMenuScreen()

	MarginAttribute("0").Apply(&DocumentBody)
	DocumentBody.AppendChildren(gameScreen, menuScreen)

	menuScreen = menuScreen.Replaced(createMenuScreen())

	<-make(chan struct{})
}

func beginGame(gameID string) {
	gameScreen.Replaced(createGameScreen(gameID))
	menuScreen = menuScreen.ReplacedWithEmpty()
}
