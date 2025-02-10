package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		count  int
		number string
	)

	fmt.Fscan(in, &count)
	for count > 0 {
		fmt.Fscan(in, &number)

		length := len(number)
		if length < 2 {
			fmt.Fprint(out, "0\n")
		}

		// переводим первый символ в число и запоминаем
		last, err := strconv.Atoi(string(number[0]))
		if err != nil {
			panic(err)
		}

		// перебирам символы в строке, начиная со второго символа
		for i := 1; i < length; i++ {
			current, err := strconv.Atoi(string(number[i]))
			if err != nil {
				panic(err)
			}

			// если текущий элемент больше предыдущего
			if last < current {
				fmt.Fprint(out, number[0:i-1], number[i:], "\n")
				break
			}
			// если это был последний элемент и не было найдено нужного числа
			if i+1 == length {
				fmt.Fprint(out, number[0:i], "\n")
				break
			}
			last = current
		}

		count--
	}
}
