//go:build wasm && js

package web

import (
	. "github.com/aliics/hearts/web/dom"
)

type screenChange struct {
	screenType
	value string
}

type screenType uint8

const (
	screenMenu screenType = iota
	screenGame
)

var screenChangeCh = make(chan screenChange)

func Run() {
	var gameScreen Element
	menuScreen := createMenuScreen()
	DocumentBody.AppendChildren(menuScreen)

	for {
		switch c := <-screenChangeCh; c.screenType {
		case screenMenu:
			DocumentBody.RemoveChild(gameScreen)
			DocumentBody.AppendChild(menuScreen)
		case screenGame:
			gameScreen = createGameScreen(c.value)
			DocumentBody.RemoveChildren(menuScreen)
			DocumentBody.AppendChild(gameScreen)
		}
	}
}

func beginGame(gameID string) {
	screenChangeCh <- screenChange{screenGame, gameID}
}
