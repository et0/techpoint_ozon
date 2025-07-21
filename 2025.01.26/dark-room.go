package main

import (
	"bufio"
	"fmt"
	"os"
)

type lamp struct {
	x       int
	y       int
	dirrect string
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		t    int // кол-во наборов
		n, m int // размер комнаты
	)

	fmt.Fscan(in, &t)

	for t > 0 {
		fmt.Fscan(in, &n, &m)

		lamps := make([]lamp, 0, 100)

		for i, j := 1, 1; i*j <= m*n; {
			// Светим в большую сторону
			if n < m {
				lamps = append(lamps, lamp{i, j, "R"})

				if i >= n {
					break
				}
				i, j = n, n-1
				lamps = append(lamps, lamp{i, j, "L"})
			} else {
				lamps = append(lamps, lamp{i, j, "D"})

				if j >= m {
					break
				}
				i, j = m-1, m
				lamps = append(lamps, lamp{i, j, "U"})
			}

			break
		}

		fmt.Fprintln(out, len(lamps))
		for _, v := range lamps {
			fmt.Fprintln(out, v.x, v.y, v.dirrect)
		}

		t--
	}
}
