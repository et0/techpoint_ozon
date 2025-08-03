package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func draw(grid *[][]byte, width, height *int, offsetN, offsetM int) {
	offsetM++
	offsetN++

	offset := offsetM + *height
	// copy((*grid)[offsetN][offset:], []byte(fmt.Sprintf("%s", strings.Repeat("_", *width))))
	for i := 0; i < *width; i++ {
		(*grid)[offsetN][offset+i] = '_'
		(*grid)[offsetN+*height*2][offset+i] = '_'
	}
	offsetN++

	space := *height - 1
	for i := 0; i < *height; i++ {
		(*grid)[offsetN+i][offsetM+space] = '/'
		(*grid)[offsetN+i][offsetM+space+*width+2*i+1] = '\\'

		(*grid)[offsetN+*height*2-i-1][offsetM+space] = '\\'
		(*grid)[offsetN+*height*2-i-1][offsetM+space+*width+2*i+1] = '/'

		space--
	}
}

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

	grid := make([][]byte, n+2)

	line := []byte(fmt.Sprintf("%s%s%s", "+", strings.Repeat("-", m), "+"))
	grid[0] = line
	grid[n+1] = line

	for i := 1; i < n+1; i++ {
		grid[i] = []byte(fmt.Sprintf("%s%s%s", "|", strings.Repeat(" ", m), "|"))
	}

	var offsetX, offsetY, counterInline int
	var lastLine bool
	for k > 0 {
		draw(&grid, &width, &height, offsetX, offsetY)
		counterInline++
		k--

		if offsetY+width+height*2 < m-height*2 {
			if !lastLine {
				if counterInline%2 != 0 {
					offsetX += height
				} else {
					offsetX -= height
				}
				offsetY += width + height
			} else {
				offsetY += (width + height) * 2
			}
		} else {
			if counterInline%2 == 0 {
				offsetX += height
			} else {
				offsetX += height * 2
			}

			if offsetX+height*2+1 >= n-height-1 {
				lastLine = true
			}

			offsetY = 0
			counterInline = 0
		}
	}

	for i := 0; i < n+2; i++ {
		for j := 0; j < m+2; j++ {
			fmt.Printf("%c", grid[i][j])
		}
		fmt.Println()
	}

}
