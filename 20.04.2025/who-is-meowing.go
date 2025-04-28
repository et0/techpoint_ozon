package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		t int
		n int
	)

	re, _ := regexp.Compile(`^([A-Za-z]+)\: ([A-Za-z]+) (am\snot|is\snot|am|is) ([A-Za-z]+)!`)

	fmt.Fscan(in, &t)
	for t > 0 {
		// Мапа с подозреваемыми
		suspects := make(map[string]int)

		// действие
		var statement string

		fmt.Fscan(in, &n)
		for i := 0; i < n; {
			line, err := in.ReadString('\n')
			if err != nil {
				break
			}
			if line == "\n" {
				continue
			}

			/*
				1 - Кто говорит (Имя1)
				2 - Про кого говорит (Имя2)
				3 - вспомогательный глагол
				4 - действи
			*/
			find := re.FindStringSubmatch(line)

			// Сохраняем действие, если оно ещё не было сохранено
			if statement == "" {
				statement = find[4]
			}

			// Проверяем наличие подозреваемых в мапе
			if _, ok := suspects[find[1]]; !ok {
				suspects[find[1]] = 0
			}
			if find[2] != "I" && find[3] != "am" && find[3] != "am not" {
				if _, ok := suspects[find[2]]; !ok {
					suspects[find[2]] = 0
				}
			}

			// Проверяем высказывание
			switch find[3] {
			case "is":
				suspects[find[2]]++
			case "is not":
				suspects[find[2]]--
			case "am not":
				suspects[find[1]]--
			case "am":
				suspects[find[1]] += 2
			}

			i++
		}

		names := make([]string, 0, len(suspects))
		max := -2147483648
		for k, v := range suspects {
			if v > max {
				names = names[:0]
				names = append(names, k)
				max = v
			} else if v == max {
				names = append(names, k)
			}
		}

		// Делаем сортировку слайся, если больше одной записи
		if len(names) > 1 {
			sort.Strings(names)
		}

		for _, v := range names {
			fmt.Fprintf(out, "%s is %s.\n", v, statement)
		}

		t--
	}
}
