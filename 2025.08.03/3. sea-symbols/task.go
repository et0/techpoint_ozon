package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func findDefaultHexagon(source *[][]byte, m, n int) (int, int) {
	height, width := 0, 0

	// проходим по первой строчке, что бы найти первый символ _ и вычислить длину и высоту плитки
	for j := 0; j < m; j++ {
		if (*source)[0][j] != '_' {
			if height != 0 {
				width = j - width
				break
			}
			continue
		}

		if width == 0 {
			width = j
		} else {
			continue
		}

		for i := 1; i < n; i++ {
			if (*source)[i][j] == '_' {
				height += i
				break
			}
		}
	}

	return height, width
}

func createHexagon(height, width int) *[][]byte {
	hexagon := make([][]byte, height+1)

	// первая линия состоящая из нижнего пробела
	hexagon[0] = []byte(strings.Repeat("_", width))

	line := 1
	for i := 0; i < height/2; i++ {
		hexagon[line] = []byte(fmt.Sprintf("%s%s%s", "/", strings.Repeat(" ", width+2*i), "\\"))
		line++
	}
	for i := height/2 - 1; i > 0; i-- {
		hexagon[line] = []byte(fmt.Sprintf("%s%s%s", "\\", strings.Repeat(" ", width+2*i), "/"))
		line++
	}

	hexagon[line] = []byte(fmt.Sprintf("%s%s%s", "\\", strings.Repeat("_", width), "/"))

	return &hexagon
}

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

		source := make([][]byte, n)
		resultMap := make([][]byte, n)
		for i := 0; i < n; {
			line, err := in.ReadString('\n')
			if err != nil {
				break
			}
			if line == "\n" {
				continue
			}

			source[i] = []byte(strings.TrimRight(line, "\n"))
			resultMap[i] = []byte(fmt.Sprintf("%s", strings.Repeat("~", m)))

			i++
		}

		// находим высоту и ширину эталона
		height, width := findDefaultHexagon(&source, m, n)

		// создаём эталонную фигуру по высоте и ширине
		hexagon := createHexagon(height, width)

		// Проходим по карте, до допустимой высоты и ширины
		for i := 0; i < n-height; i++ {
			for j := 0; j < m-width; j++ {
				// первый попавшийся символ _ - считаем отсчётом начала фигуры
				if source[i][j] != '_' {
					continue
				}

				// сравниваем фигуру с оригинало
				if compare(&source, hexagon, i, j) {
					// фигуры одинаковые, копируем фигуру в конечную карту
					copyHexagon(&resultMap, hexagon, i, j)
				}

				// смещенеи по строке на длину стандартной фигуры
				j += width
			}
		}

		// Печатаем итоговую карту
		for i := 0; i < n; i++ {
			fmt.Fprintln(out, string(resultMap[i]))
		}
		fmt.Fprint(out, "\n")
	}
}

func copyHexagon(resultMap *[][]byte, hexagon *[][]byte, oi, oj int) {
	for i := 0; i < len(*hexagon); i++ {
		if i != 0 && len((*hexagon)[i-1]) != len((*hexagon)[i]) {
			if len((*hexagon)[i-1]) < len((*hexagon)[i]) {
				oj--
			} else {
				oj++
			}
		}
		copy((*resultMap)[oi+i][oj:oj+len((*hexagon)[i])], (*hexagon)[i])
	}
}

func compare(source *[][]byte, hexagon *[][]byte, si, sj int) bool {
	for i := 0; i < len(*hexagon); i++ {
		// смщение строки вправо или влево, если увечивается длина строки исходного шестиугольника
		if i != 0 && len((*hexagon)[i-1]) != len((*hexagon)[i]) {
			if len((*hexagon)[i-1]) < len((*hexagon)[i]) {
				sj--
			} else {
				sj++
			}
		}

		// сравниваем
		if !bytes.Equal((*hexagon)[i], (*source)[si+i][sj:sj+len((*hexagon)[i])]) {
			return false
		}
	}

	return true
}
