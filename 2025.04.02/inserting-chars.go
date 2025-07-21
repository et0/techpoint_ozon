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

	var t int // кол-во тестов

	fmt.Fscan(in, &t)
	for t > 0 {
		var s string

		fmt.Fscan(in, &s)
		length := len(s)

		t--

		if length == 1 { // only yes
			fmt.Fprintln(out, "YES")
			continue
		}

		if length == 2 { // gg - yes, hf - no
			if s[0] != s[1] {
				fmt.Fprintln(out, "NO")
			} else {
				fmt.Fprintln(out, "YES")
			}
			continue
		}

		if s[0] != s[length-1] { // first != last - NO
			fmt.Fprintln(out, "NO")
			continue
		}

		var flagNo bool
		for i := 1; i < length-1; i++ {
			if s[i] == s[0] || s[i+1] == s[0] {
				continue
			}
			flagNo = true
			break
		}
		if flagNo {
			fmt.Fprintln(out, "NO")
			continue
		}
		fmt.Fprintln(out, "YES")
	}
}
