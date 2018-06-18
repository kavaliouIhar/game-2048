package main

import (
	"github.com/kavaliouIhar/game-2048/board"
)

func main() {
	board.StartMessage()
	a := board.CreateBoard(4)
	a.Play()
}
