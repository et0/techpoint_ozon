package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cord struct {
	mark byte // метка пути
	n, m int  // координата
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		t, n, m, n1, m1, n2, m2 int
	)
	fmt.Fscan(in, &t)
	for t > 0 {
		t--

		fmt.Fscan(in, &n, &m)

		sourceMap := make([][]byte, n)

		for i := 0; i < n; {
			line, err := in.ReadString('\n')
			if err != nil {
				break
			}
			if line == "\n" {
				continue
			}

			sourceMap[i] = []byte(strings.TrimRight(line, "\n"))
			// resultMap[i] = []byte(fmt.Sprintf("%s", strings.Repeat("~", m)))

			i++
		}

		fmt.Fscan(in, &n1, &m1)
		fmt.Fscan(in, &n2, &m2)

		// Если координаты одинаковые
		if n1 == n2 && m1 == m2 {
			fmt.Fprintln(out, "YES")
			continue
		}

		// находим высоту и ширину эталона
		height, width := findDefaultHexagon(&sourceMap, m, n)

		// находим верхний угол у двух фигур
		startN1, startM1 := findStartHexagon(&sourceMap, n1, m1)
		startN2, startM2 := findStartHexagon(&sourceMap, n2, m2)

		// Если две координаты на одной фигуре
		if startN1 == startN2 && startM1 == startM2 {
			fmt.Fprintln(out, "YES")
			continue
		}

		// проставляем начальные метки на карту
		sourceMap[startN1][startM1] = '1'
		sourceMap[startN2][startM2] = '2'

		// Создаём слайс накопитель возможных путей движения
		// cap равно максимум, кол-ва элементов на карте
		cords := make([]*cord, 0, (m/(height+width))*(n/height+1))

		// Добавляем исходные координаты с которых начинается движение в стороны
		cords = append(cords,
			&cord{mark: '1', n: startN1, m: startM1}, // x1
			&cord{mark: '2', n: startN2, m: startM2}, // x2
		)

		if check(&sourceMap, &cords, n, m, height, width) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}

		// Печатаем итоговую карту
		for i := 0; i < n; i++ {
			fmt.Fprintln(out, string(sourceMap[i]))
		}
		fmt.Fprint(out, "\n")

		break
	}
}

func check(sourceMap *[][]byte, cords *[]*cord, n, m, height, width int) bool {
	directions := [6][2]int{
		{height * -1, (height + width) * -1}, // up left
		{height * -2, 0},                     // up
		{height * -1, height + width},        // up right
		{height, (height + width) * -1},      // down left
		{height * 2, 0},                      // down
		{height, height + width},             // down right
	}
	//fmt.Println(directions)

	for size := len(*cords); size > 0; size = len(*cords) {
		for _, c := range *cords {
			// проверяем фигуры в шести направлениях и добавляем метки
			for _, d := range directions {
				// проверка выхода за координаты
				if c.n+d[0] <= 0 || c.n+d[0] >= n-height || c.m+d[1] <= 0 || c.m+d[1] >= m {
					continue
				}

				// если это море
				if (*sourceMap)[c.n+d[0]][c.m+d[1]] == '~' {
					continue
				}

				// Если при смещение в сторону, на фигуре уже есть противоположная метка, значит пути пересекаются
				// Если метка таже, значит уже были в этом направление
				if (*sourceMap)[c.n+d[0]][c.m+d[1]] == '1' {
					if c.mark == '2' {
						return true
					}
					continue
				} else if (*sourceMap)[c.n+d[0]][c.m+d[1]] == '2' {
					if c.mark == '1' {
						return true
					}
					continue
				}

				// TODO: проверить, есть ли в этих координатах фигура
				fmt.Println("D:", d, "NM:", c.n, c.m)
				if !isHexagonByCord(sourceMap, c.n+d[0], c.m+d[1], height, width) {
					// фигуры нет, ставим символ моря
					(*sourceMap)[c.n+d[0]][c.m+d[1]] = '~'
					continue
				}

				// ставим метку соответствующего пути
				if (*sourceMap)[c.n+d[0]][c.m+d[1]] == ' ' {
					(*sourceMap)[c.n+d[0]][c.m+d[1]] = c.mark

					// добавляем новые координаты
					*cords = append(*cords, &cord{c.mark, c.n + d[0], c.m + d[1]})
				}

			}
		}

		// Смещаем проверенные элементы
		(*cords) = (*cords)[size:]
	}

	return false
}

func isHexagonByCord(sourceMap *[][]byte, x, y, height, width int) bool {
	for i := 0; i < width; i++ {
		// fmt.Println(string((*sourceMap)[x-1][y+i]), string((*sourceMap)[x+height*2-1][y+i]))
		if (*sourceMap)[x-1][y+i] != '_' || (*sourceMap)[x+height*2-1][y+i] != '_' {
			return false
		}
	}

	for i := 0; i < height; i++ {
		// fmt.Println(string((*sourceMap)[x+i][y-i-1]), string((*sourceMap)[x+i][y+width+i]))
		if (*sourceMap)[x+i][y-i-1] != '/' || (*sourceMap)[x+i][y+width+i] != '\\' {
			return false
		}

		// fmt.Println(string((*sourceMap)[x+height*2-1-i][y-i-1]), string((*sourceMap)[x+height*2-1-i][y+width+i]))
		if (*sourceMap)[x+height*2-1-i][y-i-1] != '\\' || (*sourceMap)[x+height*2-1-i][y+width+i] != '/' {
			return false
		}
	}

	return true
}

func findStartHexagon(sourceMap *[][]byte, n, m int) (int, int) {
	for ; n > 0; n-- {
		if (*sourceMap)[n-1][m] == '/' {
			m++
		} else if (*sourceMap)[n-1][m] == '\\' {
			m--
		} else if (*sourceMap)[n-1][m] == '_' {
			break
		}
	}

	for ; m > 0; m-- {
		if (*sourceMap)[n][m-1] != '/' {
			continue
		}

		break
	}

	return n, m
}

func findDefaultHexagon(sourceMap *[][]byte, m, n int) (int, int) {
	height, width := 0, -1

	// проходим по первой строчке, что бы найти первый символ _ и вычислить длину фигуры
	j := 0
	for ; j < m; j++ {
		if (*sourceMap)[0][j] != '_' {
			if width == -1 {
				continue
			}

			width = j - width

			// возвращаем начальную позицию
			j -= width

			break
		}

		if width == -1 {
			width = j
		}
	}

	// вычисляем высоту фигуры, смещяемся вниз и вправо, пока не достигнем стенки (j == 0) или пока не встретится символ /
	for i := 1; i < n; i++ {
		j--
		if j >= 0 && (*sourceMap)[i][j] == '/' {
			height++

			continue
		}

		break
	}

	return height, width
}
