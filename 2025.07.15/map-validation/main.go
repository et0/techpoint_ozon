package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--

		var (
			n, m int
		)
		fmt.Fscan(in, &n, &m)

		counter := make(map[byte]int)
		startCoords := make(map[byte][2]int)

		visited := make([][]bool, n)
		fields := make([][]byte, n)
		for i := 0; i < n; i++ {
			fields[i] = make([]byte, m)
			visited[i] = make([]bool, m)
		}

		for i := 0; i < n; i++ {
			line := make([]byte, m)
			fmt.Fscan(in, &line)

			for j, v := range line {
				if v == '.' {
					continue
				}

				counter[v]++
				if counter[v] == 1 {
					startCoords[v] = [2]int{i, j}
				}

				fields[i][j] = v
			}
		}

		var errorFlag bool
		for i, coord := range startCoords {
			check(i, coord[0], coord[1], n, m, &fields, &counter, &visited)
			if counter[i] != 0 {
				errorFlag = true
				break
			}
		}

		if errorFlag {
			fmt.Fprintln(out, "NO")
		} else {
			fmt.Fprintln(out, "YES")
		}
	}
}

func check(key byte, x, y, maxN, maxM int, fields *[][]byte, counter *map[byte]int, visited *[][]bool) {
	if (*fields)[x][y] != key {
		return
	}
	(*counter)[key]--
	(*visited)[x][y] = true

	// если справа есть элементы
	if y+2 < maxM && !(*visited)[x][y+2] {
		check(key, x, y+2, maxN, maxM, fields, counter, visited)
	}
	// если слева есть элементы
	if y-2 >= 0 && !(*visited)[x][y-2] {
		check(key, x, y-2, maxN, maxM, fields, counter, visited)
	}

	if x > 0 {
		// если справа есть элементы
		if y+1 < maxM && !(*visited)[x-1][y+1] {
			check(key, x-1, y+1, maxN, maxM, fields, counter, visited)
		}
		// если слева есть элементы
		if y-1 >= 0 && !(*visited)[x-1][y-1] {
			check(key, x-1, y-1, maxN, maxM, fields, counter, visited)
		}
	}

	// это последняя строка
	if x+1 >= maxN {
		return
	}

	// если справа есть элементы
	if y+1 < maxM && !(*visited)[x+1][y+1] {
		check(key, x+1, y+1, maxN, maxM, fields, counter, visited)
	}
	// если слева есть элементы
	if y-1 >= 0 && !(*visited)[x+1][y-1] {
		check(key, x+1, y-1, maxN, maxM, fields, counter, visited)
	}
}
