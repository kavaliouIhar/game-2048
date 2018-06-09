package board

import "fmt"

func (b *Board) SlideLeft() {
	if !b.CheckLeft() {
		fmt.Println("can't slide left")
		return
	}
	q := b.size
	for i := 0; i < q; i++ {
		t := b.GetRow(i)
		for k, numx := 0, 0; k < q; k++ {
			if t[k] != 0 {
				if numx != k {
					t[numx], t[k] = t[k], 0
				}
				numx++
			}
		}
		for j := 0; j < q-1; j++ {
			if t[j] != 0 && t[j] == t[j+1] {
				t[j], t[j+1] = t[j]*2, 0
				b.score += t[j]
			}
		}
		for k, numx := 0, 0; k < q; k++ {
			if t[k] != 0 {
				if numx != k {
					t[numx], t[k] = t[k], 0
				}
				numx++
			}
		}
		b.SetRow(i, t)
	}
	b.AutoInsert()
	b.Show()
}

func (b *Board) SlideUp() {
	if !b.CheckUp() {
		fmt.Println("can't slide up")
		return
	}
	q := b.size
	for i := 0; i < q; i++ {
		t := b.GetCol(i)
		for k, numx := 0, 0; k < q; k++ {
			if t[k] != 0 {
				if numx != k {
					t[numx], t[k] = t[k], 0
				}
				numx++
			}
		}
		for j := 0; j < q-1; j++ {
			if t[j] != 0 && t[j] == t[j+1] {
				t[j], t[j+1] = t[j]*2, 0
				b.score += t[j]
			}
		}
		for k, numx := 0, 0; k < q; k++ {
			if t[k] != 0 {
				if numx != k {
					t[numx], t[k] = t[k], 0
				}
				numx++
			}
		}
		b.SetCol(i, t)
	}
	b.AutoInsert()
	b.Show()
}

func (b *Board) SlideRight() {
	if !b.CheckRight() {
		fmt.Println("can't slide right")
		return
	}
	q := b.size
	for i := 0; i < q; i++ {
		t := b.GetRow(i)
		for k, numx := q-1, q-1; k >= 0; k-- {
			if t[k] != 0 {
				if numx != k {
					t[numx], t[k] = t[k], 0
				}
				numx--
			}
		}
		for j := q - 1; j > 0; j-- {
			if t[j] != 0 && t[j] == t[j-1] {
				t[j], t[j-1] = t[j]*2, 0
				b.score += t[j]
			}
		}
		for k, numx := q-1, q-1; k >= 0; k-- {
			if t[k] != 0 {
				if numx != k {
					t[numx], t[k] = t[k], 0
				}
				numx--
			}
		}
		b.SetRow(i, t)
	}
	b.AutoInsert()
	b.Show()
}

func (b *Board) SlideDown() {
	if !b.CheckDown() {
		fmt.Println("can't slide down")
		return
	}
	q := b.size
	for i := 0; i < q; i++ {
		t := b.GetCol(i)
		for k, numx := q-1, q-1; k >= 0; k-- {
			if t[k] != 0 {
				if numx != k {
					t[numx], t[k] = t[k], 0
				}
				numx--
			}
		}
		for j := q - 1; j > 0; j-- {
			if t[j] != 0 && t[j] == t[j-1] {
				t[j], t[j-1] = t[j]*2, 0
				b.score += t[j]
			}
		}
		for k, numx := q-1, q-1; k >= 0; k-- {
			if t[k] != 0 {
				if numx != k {
					t[numx], t[k] = t[k], 0
				}
				numx--
			}
		}
		b.SetCol(i, t)
	}
	b.AutoInsert()
	b.Show()
}
