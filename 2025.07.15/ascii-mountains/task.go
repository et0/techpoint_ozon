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
			k, n, m int
		)
		fmt.Fscan(in, &k, &n, &m)

		mounting := make([][]byte, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &mounting[i])
		}
		k--

		for k > 0 {
			k--

			for i := 0; i < n; i++ {
				line := make([]byte, m)
				fmt.Fscan(in, &line)

				for k, v := range line {
					if mounting[i][k] != '.' {
						continue
					}
					mounting[i][k] = v
				}
			}
		}

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Fprint(out, string(mounting[i][j]))
			}
			fmt.Fprint(out, "\n")
		}
		fmt.Fprint(out, "\n")
	}
}
