package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	name byte
	n    int
	m    int
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		t int // кол-во наборов
		n int // кол-во строк
		m int // кол-во столбов
	)

	maps := map[byte]string{
		46: ".",
		35: "#",
		65: "A",
		66: "B",
		97: "a",
		98: "b",
	}

	fmt.Fscan(in, &t)
	// fmt.Fprintln(out, "T: ", t)

	for t > 0 {
		fmt.Fscan(in, &n, &m)
		// fmt.Fprintln(out, "N:", n, "M:", m)

		field := make([][]byte, n, n)    // поле
		robots := make([]position, 0, 2) // роботы и их координаты

		for x := 0; x < n; x++ {
			var line string
			fmt.Fscan(in, &line)
			field[x] = make([]byte, m, m)

			for i := 0; i < m; i++ {
				field[x][i] = byte(line[i])

				// Ищем в строке А или В
				if field[x][i] == 65 || field[x][i] == 66 {
					robots = append(robots, position{field[x][i], x, i})
				}
			}
		}

		// fmt.Fprintln(out, robots)

		// Выполняем Смещение для первого робота
		// Проверка, может ли робот двигаться наверх, если нет, то смещаемся влево на одну позицию
		if robots[0].m > 0 && robots[0].m%2 != 0 {
			robots[0].m--
			field[robots[0].n][robots[0].m] = robots[0].name + 32
		}
		// Смещение вверх до первой строки
		for robots[0].n > 0 {
			robots[0].n--
			field[robots[0].n][robots[0].m] = robots[0].name + 32
		}
		// Смещение влево до первого столбца
		for robots[0].m > 0 {
			robots[0].m--
			field[robots[0].n][robots[0].m] = robots[0].name + 32
		}

		// Выполняем смещение для второго робота
		// Проверка, может ли робот двигаться вниз, если нет, то смещаемся вправо на одну позицию
		if robots[1].m < m-1 && robots[1].m%2 != 0 {
			robots[1].m++
			field[robots[1].n][robots[1].m] = robots[1].name + 32
		}
		// Смещение вниз до последней строки
		for robots[1].n < n-1 {
			robots[1].n++
			field[robots[1].n][robots[1].m] = robots[1].name + 32
		}
		// Смещение вправо до последнего столбца
		for robots[1].m < m-1 {
			robots[1].m++
			field[robots[1].n][robots[1].m] = robots[1].name + 32
		}

		for x := 0; x < n; x++ {
			for y := 0; y < m; y++ {
				fmt.Fprint(out, maps[field[x][y]])
			}
			fmt.Fprintln(out, "")
		}
		//fmt.Fprintln(out, robots, "\n")

		t--
	}
}
