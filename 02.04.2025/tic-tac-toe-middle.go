package main

import (
	"bufio"
	"fmt"
	"os"
)

type direction struct {
	val   byte
	X     int  // repeat X in direction
	O     int  // repeat 0 in direction
	empty bool // has empty ceil
}

type line struct {
	left   direction
	middle direction
	right  direction
}

// Условие победы X, если поставить один X
func canXWin(line *direction, k *int) bool {
	//fmt.Println("X ", line)
	if line.X == *k-1 && line.empty {
		return true
	}
	return false
}

// Условие победы или X или O
func canAnotherWin(line *direction, k *int) bool {
	//fmt.Println("A ", line)
	if line.X == *k || line.O == *k {
		return true
	}

	return false
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		t    int // кол-во тестов
		k    int
		n, m int
	)

	fmt.Fscan(in, &t)
	for t > 0 {
		t--

		fmt.Fscan(in, &k, &n, &m)

		// слайс хранит накопленный данные в трех направлениях предыдущей строки
		prev := make([]line, m)
		var (
			winX       bool // может ли Х победить за один ход?
			winAnother bool // победить нельзя
			lineNumber int  // номер строки
		)
		for lineNumber < n {
			var s []byte
			fmt.Fscan(in, &s)

			//
			left := direction{}

			fmt.Println(prev)
			// копия
			oldLeft := prev[0].left
			for i, v := range s {
				// left
				if i > 0 {
					// копия
					oldLeft = prev[i].left

					switch v {
					case 'X':
						prev[i].left.X = oldLeft.X + 1
						prev[i].left.O = 0
						prev[i].left.empty = oldLeft.empty
					case 'O':
						prev[i].left.O = oldLeft.O + 1
						prev[i].left.X = 0
						prev[i].left.empty = oldLeft.empty
					case '.':
						if oldLeft.empty {
							if oldLeft.val == 'X' {
								prev[i].left.X, prev[i].left.O = 1, 0
							} else if oldLeft.val == 'O' {
								prev[i].left.X, prev[i].left.O = 0, 1
							} else {
								prev[i].left.X, prev[i].left.O = 0, 0
							}
						} else {
							prev[i].left.empty = true
						}
					}
					prev[i].left.val = v
				}

				// middle
				switch v {
				case 'X':
					prev[i].middle.X++
					prev[i].middle.O = 0
				case 'O':
					prev[i].middle.O++
					prev[i].middle.X = 0
				case '.':
					if prev[i].middle.empty {
						if prev[i].middle.val == 'X' {
							prev[i].middle.X, prev[i].middle.O = 1, 0
						} else if prev[i].middle.val == 'O' {
							prev[i].middle.X, prev[i].middle.O = 0, 1
						} else {
							prev[i].middle.X, prev[i].middle.O = 0, 0
						}
					} else {
						prev[i].middle.empty = true
					}
				}
				prev[i].middle.val = v

				// right
				if i+1 < m {
					switch v {
					case 'X':
						prev[i].right.X = prev[i+1].right.X + 1
						prev[i].right.O = 0
						prev[i].right.empty = prev[i+1].right.empty
					case 'O':
						prev[i].right.O = prev[i].right.O + 1
						prev[i].right.X = 0
						prev[i].right.empty = prev[i+1].right.empty
					case '.':
						if prev[i+1].right.empty {
							if prev[i+1].right.val == 'X' {
								prev[i].right.X, prev[i].right.O = 1, 0
							} else if prev[i+1].right.val == 'O' {
								prev[i].right.X, prev[i].right.O = 0, 1
							} else {
								prev[i].right.X, prev[i].right.O = 0, 0
							}
						} else {
							prev[i].right.empty = true
						}
					}
					prev[i].right.val = v
				}

				if i+1 < m {
					// prev[i+1].left = prev[i].middle
					switch v {
					case 'X':
						prev[i+1].left.X = prev[i].left.X + 1
						prev[i+1].left.O = 0
						prev[i+1].left.empty = prev[i].left.empty
					case 'O':
						prev[i+1].left.O = prev[i].left.O + 1
						prev[i+1].left.X = 0
						prev[i+1].left.empty = prev[i].left.empty
					case '.':
						if prev[i].left.empty {
							if prev[i].left.val == 'X' {
								prev[i+1].left.X, prev[i+1].left.O = 1, 0
							} else if prev[i].left.val == 'O' {
								prev[i+1].left.X, prev[i+1].left.O = 0, 1
							} else {
								prev[i+1].left.X, prev[i+1].left.O = 0, 0
							}
						} else {
							prev[i+1].left.empty = true
						}
					}
					prev[i+1].left.val = v
				}
				if i > 0 {
					switch v {
					case 'X':
						prev[i-1].right.X = prev[i].right.X + 1
						prev[i-1].right.O = 0
						prev[i-1].right.empty = prev[i].right.empty
					case 'O':
						prev[i-1].right.O = prev[i].right.O + 1
						prev[i-1].right.X = 0
						prev[i-1].right.empty = prev[i].right.empty
					case '.':
						if prev[i].right.empty {
							if prev[i].right.val == 'X' {
								prev[i-1].right.X, prev[i-1].right.O = 1, 0
							} else if prev[i].right.val == 'O' {
								prev[i-1].right.X, prev[i-1].right.O = 0, 1
							} else {
								prev[i-1].right.X, prev[i-1].right.O = 0, 0
							}
						} else {
							prev[i-1].right.empty = true
						}
					}
					prev[i-1].right.val = v
				}

				// проверка слева
				switch v {
				case 'X':
					left.X++
					left.O = 0
				case 'O':
					left.O++
					left.X = 0
				case '.':
					if left.empty { // если уже был пробел ранее
						if s[i-1] == '.' {
							left.X = 0
						} else if left.X > 0 {
							left.X = 1
						}
					} else {
						left.empty = true
					}
					left.O = 0
				}

				// Проверка, может ли победить X
				if !winX {
					winX = canXWin(&left, &k) || canXWin(&prev[i].middle, &k)
				}
				// Победил ли кто-то уже
				winAnother = canAnotherWin(&left, &k) || canAnotherWin(&prev[i].middle, &k)
				if winAnother {
					break
				}
			}
			fmt.Println(prev)

			if winAnother {
				break
			}

			lineNumber++
		}

		if winX && !winAnother {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
		break
	}
}
