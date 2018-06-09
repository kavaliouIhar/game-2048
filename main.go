package main

import (
	"myprojects/game-2048/board"
)

func main() {
	board.StartMessage()
	a := board.CreateBoard(4)
	for {
		if a.GameOver() {
			a.Loss()
		} else {
			a.Play()
		}
	}
}
