package board

func (b *Board) CheckUp() bool {
	q := b.size
	for i := q - 1; i > 0; i-- {
		for j := q - 1; j >= 0; j-- {
			dex := i*q + j
			if b.data[dex] != 0 {
				if b.data[dex-q] == 0 {
					return true
				}
				if b.data[dex-q] == b.data[dex] {
					return true
				}
			}

		}
	}
	return false
}

func (b *Board) CheckLeft() bool {
	q := b.size
	for i := 0; i < q; i++ {
		t := b.data[i*q : (i+1)*q]
		for k := 1; k < q; k++ {
			if t[k] != 0 {
				if t[k-1] == 0 {
					return true
				}
				if t[k-1] == t[k] {
					return true
				}
			}

		}
	}
	return false
}

func (b *Board) CheckRight() bool {
	q := b.size
	for i := 0; i < q; i++ {
		t := b.data[i*q : (i+1)*q]
		for k := 0; k < q-1; k++ {
			if t[k] != 0 {
				if t[k+1] == 0 {
					return true
				}
				if t[k+1] == t[k] {
					return true
				}
			}

		}
	}
	return false
}

func (b *Board) CheckDown() bool {
	q := b.size
	for i := q - 2; i >= 0; i-- {
		for j := q - 1; j >= 0; j-- {
			dex := i*q + j
			if b.data[dex] != 0 {
				if b.data[dex+q] == 0 {
					return true
				}
				if b.data[dex+q] == b.data[dex] {
					return true
				}
			}

		}
	}
	return false
}
