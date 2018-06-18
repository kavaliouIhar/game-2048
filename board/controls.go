package board

import (
	"fmt"
	"os"
)

func (b *Board) End() {
	b.Show()
	fmt.Println("u lost")
	fmt.Println("ur final score is: ", b.score)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("enter r to restart or enter e to exit")
	fmt.Print("Enter command: ")
	var input string
	fmt.Scanln(&input)
	if input == "r" {
		b.RestartBoard()
	} else if input == "e" {
		os.Exit(2)
	} else {
		fmt.Println("wrong command")
	}

}

func (b *Board) Continue() {
	fmt.Print("Enter command: ")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "w":
		b.SlideUp()
	case "a":
		b.SlideLeft()
	case "s":
		b.SlideDown()
	case "d":
		b.SlideRight()
	case "r":
		b.RestartBoard()
	case "e":
		os.Exit(2)
	default:
		fmt.Println("wrong command")
	}
}

func (b *Board) Play() {
	for {
		if b.GameOver() {
			b.End()
		} else {
			b.Continue()
		}
	}
}
