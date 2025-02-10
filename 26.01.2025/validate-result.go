package main

import (
	"bufio"
	"fmt"
	"os"
)

type tmpl struct {
	key   string
	value string
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

	fmt.Fscan(in, &t)
	for t > 0 {
		fmt.Fscan(in, &n)

		data := make(map[string]map[string]bool)
		for i := 0; i < n; i++ {
			var name, price string

			fmt.Fscan(in, &name, &price)
			if _, ok := data[price]; !ok {
				data[price] = make(map[string]bool)
			}
			data[price][name] = true
		}

		var result string
		fmt.Fscan(in, &result)

		//fmt.Fprintln(out, 100001-t, data, result)

		element := tmpl{}
		index := struct {
			start  int // позиция начала нового элемента в строке результата
			status int // флаг перехода, если 0 - читается название до ":" ; если 1 - читается кол-во до ","
		}{0, 0}
		flag := true
		for i := 0; i < len(result); i++ {
			if string(result[i]) == ":" {
				if index.status == 1 {
					flag = false
					break
				}
				element.key = string(result[index.start:i])
				index.start = i + 1
				index.status = 1
			}
			if string(result[i]) == "," || i == len(result)-1 {
				if i == len(result)-1 {
					element.value = string(result[index.start:len(result)])
				} else {
					element.value = string(result[index.start:i])
					index.start = i + 1
					index.status = 0
				}

				//fmt.Fprintln(out, i, data, element)

				if element.key == "" || element.value == "" {
					flag = false
					break
				}
				if _, ok := data[element.value]; !ok {
					flag = false
					break
				} else if _, ok := data[element.value][element.key]; !ok {
					flag = false
					break
				}
				delete(data, element.value)
				element = tmpl{}
			}
		}

		if !flag || len(data) > 0 {
			fmt.Fprintln(out, "NO")
		} else {
			fmt.Fprintln(out, "YES")
		}

		t--
	}
}
