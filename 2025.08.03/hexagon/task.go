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

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--

		var (
			width, height int
		)
		fmt.Fscan(in, &width, &height)

		fmt.Fprintf(out, "%s%s\n", strings.Repeat(" ", height), strings.Repeat("_", width))

		space := height - 1
		for i := 0; i < height; i++ {
			fmt.Fprintf(out, "%s%s%s%s\n", strings.Repeat(" ", space), "/", strings.Repeat(" ", width+2*i), "\\")
			space--
		}

		space = 0
		for i := height - 1; i > 0; i-- {
			fmt.Fprintf(out, "%s%s%s%s\n", strings.Repeat(" ", space), "\\", strings.Repeat(" ", width+2*i), "/")
			space++
		}

		fmt.Fprintf(out, "%s%s%s%s\n", strings.Repeat(" ", height-1), "\\", strings.Repeat("_", width), "/")

	}
}
