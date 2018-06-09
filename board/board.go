package board

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Board struct {
	size  int
	data  []int
	score int
}

func StartMessage() {
	fmt.Println("hi, this is simple 2048 game")
	fmt.Println("controls are wasd + enter")
	fmt.Println("to restart game enter r")
	fmt.Println("to exit enter e")
}

func (b *Board) RestartBoard() {
	fmt.Println("restarting game")
	b.data = make([]int, b.size*b.size)
	b.score = 0
	b.AutoInsert()
	b.AutoInsert()
	b.Show()
}

func CreateBoard(a int) *Board {
	rand.Seed(time.Now().UTC().UnixNano())

	if a < 2 || a > 10 {
		a = 4
	}
	s := &Board{size: a, data: make([]int, a*a)}
	s.AutoInsert()
	s.AutoInsert()
	s.Show()
	return s
}

func (b *Board) GetCol(a int) []int {
	q := b.size
	t := make([]int, q)
	for i := 0; i < q; i++ {
		t[i] = b.data[i*q+a]
	}
	return t
}

func (b *Board) GetRow(a int) []int {
	q := b.size
	t := make([]int, q)
	for i := 0; i < q; i++ {
		t[i] = b.data[a*q+i]
	}
	return t
}

func (b *Board) SetCol(a int, t []int) {
	q := b.size
	if a < 0 || a >= q {
		fmt.Println(errors.New("column index out of range"))
		return
	}
	if len(t) != q {
		fmt.Println(errors.New("wrong slice length"))
		return
	}
	for i := 0; i < q; i++ {
		b.data[i*q+a] = t[i]
	}
}

func (b *Board) SetRow(a int, t []int) {
	q := b.size
	if a < 0 || a >= q {
		fmt.Println(errors.New("row index out of range"))
		return
	}
	if len(t) != q {
		fmt.Println(errors.New("wrong slice length"))
		return
	}
	for i := 0; i < q; i++ {
		b.data[a*q+i] = t[i]
	}
}

func (b *Board) Show() {
	q := b.size
	fmt.Println()
	for i := 0; i < q; i++ {
		for j := 0; j < q; j++ {
			fmt.Print("|‾‾‾‾‾‾|")
		}
		fmt.Println()
		for j := 0; j < q; j++ {
			fmt.Print(showGrid(b.data[i*q+j]))
		}
		fmt.Println()
		for j := 0; j < q; j++ {
			fmt.Print("|______|")
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("score: ", b.score)
	fmt.Println()
}

func (b *Board) AutoInsert() {
	q := b.size * b.size
	r := []int{}
	for i := 0; i < q; i++ {
		if b.data[i] == 0 {
			r = append(r, i)
		}
	}
	if len(r) == 0 {
		fmt.Println("no empty slots")
	}
	x := rand.Intn(len(r))

	if k := rand.Intn(11); k < 10 {
		b.data[r[x]] = 2
	} else {
		b.data[r[x]] = 4
	}
}

func (b *Board) GameOver() bool {
	if b.CheckDown() {
		return false
	}
	if b.CheckUp() {
		return false
	}
	if b.CheckLeft() {
		return false
	}
	if b.CheckRight() {
		return false
	}
	return true
}

func showGrid(n int) string {
	if n == 0 {
		return "|      |"
	}
	if n == 2 {
		return "|    2 |"
	}
	if n == 4 {
		return "|    4 |"
	}
	if n == 8 {
		return "|    8 |"
	}
	if n == 16 {
		return "|   16 |"
	}
	if n == 32 {
		return "|   32 |"
	}
	if n == 64 {
		return "|   64 |"
	}
	if n == 128 {
		return "|  128 |"
	}
	if n == 256 {
		return "|  256 |"
	}
	if n == 512 {
		return "|  512 |"
	}
	if n == 1024 {
		return "| 1024 |"
	}
	if n == 2048 {
		return "| 2048 |"
	}
	return ""
}
