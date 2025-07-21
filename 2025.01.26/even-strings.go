package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func getSub(str string, first int, length int) string {
	newStr := make([]byte, 0, length)

	for i := first; i < length; i += 2 {
		newStr = append(newStr, str[i])
	}

	return string(newStr)
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		t int // кол-во наборов
		n int // кол-во элементов в массиве
	)

	type about struct {
		index  int
		value  string
		length int
		sub0   string
		sub1   string
	}

	fmt.Fscan(in, &t)
	for t > 0 {
		fmt.Fscan(in, &n)

		data := make([]about, n)
		//repeat0 := make(map[string]int)
		//repeat1 := make(map[string]int)
		for i := 0; i < n; i++ {
			data[i] = about{}
			fmt.Fscan(in, &data[i].value)
			data[i].length = len(data[i].value)
			data[i].sub0 = getSub(data[i].value, 0, data[i].length)
			data[i].sub1 = getSub(data[i].value, 1, data[i].length)

			/*if _, ok := repeat0[data[i].sub0]; ok {
				repeat0[data[i].sub0]++
			} else if data[i].sub0 != "" {
				repeat0[data[i].sub0] = 1
			}
			if _, ok := repeat1[data[i].sub1]; ok {
				repeat1[data[i].sub1]++
			} else if data[i].sub1 != "" {
				repeat1[data[i].sub1] = 1
			}*/
		}
		//fmt.Println(repeat0, repeat1)

		sort.Slice(data, func(i, j int) bool {
			return data[i].length < data[j].length
		})

		match := 0
		for k, d := range data {
			for i := k + 1; i < len(data); i++ {
				//fmt.Fprintln(out, "0", d.sub0, data[i].sub0[:len(d.sub0)])
				//fmt.Fprintln(out, "1", d.sub1, data[i].sub1[:len(d.sub1)])
				if d.sub0 == data[i].sub0 || (len(d.sub1) > 0 && d.sub1 == data[i].sub1) {
					match++
				}
			}
		}

		//fmt.Fprintln(out, data, match, "\n")
		fmt.Fprintln(out, match)

		t--
	}
}
