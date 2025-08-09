package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
		startN1, startM1 := findStartHexagon(&sourceMap, n1-1, m1-1)
		startN2, startM2 := findStartHexagon(&sourceMap, n2-1, m2-1)

		// Если две координаты на одной фигуре
		if startN1 == startN2 && startM1 == startM2 {
			fmt.Fprintln(out, "YES")
			continue
		}

		// проставляем начальные метки на карту
		sourceMap[startN1][startM1] = '1'
		sourceMap[startN2][startM2] = '2'

		// Создаём слайсы - накопитель возможных путей движения для двух начальных точек
		// cap равно максимальному кол-ву элементов на карте
		cords1 := make([][2]int, 0, (m/(height+width))*(n/height+1))
		cords2 := make([][2]int, 0, cap(cords1))

		// Добавляем исходные координаты с которых начинается движение в стороны
		cords1 = append(cords1, [2]int{startN1, startM1})
		cords2 = append(cords2, [2]int{startN2, startM2})

		if check(&sourceMap, &cords1, &cords2, n, m, height, width) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}

		// Печатаем итоговую карту
		// for i := 0; i < n; i++ {
		// 	fmt.Fprintln(out, string(sourceMap[i]))
		// }
		// fmt.Fprint(out, "\n")
	}
}

func check(sourceMap *[][]byte, cords1, cords2 *[][2]int, n, m, height, width int) bool {
	// шесть направление в разные стороны
	directions := [6][2]int{
		{height * -1, (height + width) * -1}, // up left
		{height * -2, 0},                     // up
		{height * -1, height + width},        // up right
		{height, (height + width) * -1},      // down left
		{height * 2, 0},                      // down
		{height, height + width},             // down right
	}
	//fmt.Println(directions)

	for size1, size2 := len(*cords1), len(*cords2); size1 > 0 && size2 > 0; size1, size2 = len(*cords1), len(*cords2) {
		// проверяем накопитель для первой точки
		for _, c := range *cords1 {
			if checkDirection(sourceMap, cords1, &c, &directions, n, m, height, width, '1') {
				return true
			}
		}

		// проверяем накопитель для второй точки
		for _, c := range *cords2 {
			if checkDirection(sourceMap, cords2, &c, &directions, n, m, height, width, '2') {
				return true
			}
		}

		// Смещаем проверенные координаты
		(*cords1) = (*cords1)[size1:]
		(*cords2) = (*cords2)[size2:]
	}

	return false
}

func checkDirection(sourceMap *[][]byte, cords *[][2]int, c *[2]int, directions *[6][2]int, n, m, height, width int, mark byte) bool {
	// проверяем фигуры в шести направлениях и добавляем метки
	for _, d := range directions {
		// проверка выхода за координаты
		if c[0]+d[0] <= 0 || c[0]+d[0] >= n-height || c[1]+d[1] <= 0 || c[1]+d[1] >= m {
			continue
		}

		// если это море
		if (*sourceMap)[c[0]+d[0]][c[1]+d[1]] == '~' {
			continue
		}

		// Если при смещение в сторону, на фигуре уже есть противоположная метка, значит пути пересекаются
		// Если метка таже, значит уже были в этом направление
		if (*sourceMap)[c[0]+d[0]][c[1]+d[1]] == '1' {
			if mark == '2' {
				return true
			}
			continue
		} else if (*sourceMap)[c[0]+d[0]][c[1]+d[1]] == '2' {
			if mark == '1' {
				return true
			}
			continue
		}

		// проверяет, есть ли в этих координатах фигура
		// fmt.Println("D:", d, "NM:", c[0], c[1])
		if !isHexagonByCord(sourceMap, c[0]+d[0], c[1]+d[1], height, width) {
			// фигуры нет, ставим символ моря
			(*sourceMap)[c[0]+d[0]][c[1]+d[1]] = '~'
			continue
		}

		// ставим метку соответствующего пути
		if (*sourceMap)[c[0]+d[0]][c[1]+d[1]] == ' ' {
			(*sourceMap)[c[0]+d[0]][c[1]+d[1]] = mark

			// добавляем новые координаты
			*cords = append(*cords, [2]int{c[0] + d[0], c[1] + d[1]})
		}
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
