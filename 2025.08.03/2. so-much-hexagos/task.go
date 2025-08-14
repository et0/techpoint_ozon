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
		m, n, width, height, k int
	)
	fmt.Fscan(in, &m, &n, &width, &height, &k)

	grid := make([][]byte, n)

	for i := 0; i < n; i++ {
		grid[i] = []byte(strings.Repeat(" ", m))
	}

	// три направления смещения, куда можно поставить фигуру
	directions := [][2]int{
		{height * -1, width + height}, // up right
		{height, width + height},      // down right)
		{0, width*2 + height*2},       // right right
	}

	// Смещение
	offsetN, offsetM := 0, height
	for k > 0 {
		draw(&grid, width, height, offsetN, offsetM)
		nextOffset(&grid, &directions, n, m, width, height, &offsetN, &offsetM)

		k--
	}

	print(out, &grid, n, m)
}

func nextOffset(sourceMap *[][]byte, directions *[][2]int, n, m, width, height int, offsetN, offsetM *int) {
	for _, d := range *directions {
		if *offsetN+d[0] < 0 {
			continue
		}

		if *offsetN+d[0] < n-height*2 && *offsetM+d[1] <= m-width-height && (*sourceMap)[*offsetN+d[0]+height*2][*offsetM+d[1]] != '_' {
			*offsetN += d[0]
			*offsetM += d[1]

			return
		}
	}

	// если направления не подошли, значит на линии закончилось место
	// следующую фигуру рисуем снизу
	*offsetM = height
	if (*sourceMap)[*offsetN+height][*offsetM] == '_' {
		*offsetN += height
	} else {
		*offsetN += height * 2
	}

}

func draw(sourceMap *[][]byte, width, height, offsetN, offsetM int) {
	for i := range width {
		(*sourceMap)[offsetN][offsetM+i] = '_'
		(*sourceMap)[offsetN+height*2][offsetM+i] = '_'
	}

	for i := range height {
		(*sourceMap)[offsetN+i+1][offsetM-i-1] = '/'
		(*sourceMap)[offsetN+i+1][offsetM+width+i] = '\\'

		(*sourceMap)[offsetN+height*2-i][offsetM-i-1] = '\\'
		(*sourceMap)[offsetN+height*2-i][offsetM+width+i] = '/'
	}
}

func print(out *bufio.Writer, sourceMap *[][]byte, n, m int) {
	fmt.Fprintf(out, "%s%s%s\n", "+", strings.Repeat("-", m), "+")
	for i := 0; i < n; i++ {
		fmt.Fprintf(out, "%s%s%s\n", "|", string((*sourceMap)[i]), "|")
	}
	fmt.Fprintf(out, "%s%s%s\n", "+", strings.Repeat("-", m), "+")
}
