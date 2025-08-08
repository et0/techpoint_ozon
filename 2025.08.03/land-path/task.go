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

		// Если координаты на одной фигуре

		// находим высоту и ширину эталона
		height, width := findDefaultHexagon(&sourceMap, m, n)
		fmt.Println(height, width)

		// TODO: попытка пойти на встречу

		break
	}
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
