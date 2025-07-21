package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type assert struct {
	who  string
	old  int
	whom string
}

func Slove() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		t int
	)

	reQuestion, _ := regexp.Compile(`^How old is ([A-Za-z]{1,16})\?\n`)
	reOne, _ := regexp.Compile(`^([A-Za-z]{1,16})\sis\s([\d]+)\syears\sold\n`)
	reTwo, _ := regexp.Compile(`^([A-Za-z]{1,16}) is the same age as ([^\s]{1,16})\n`)
	reThreeFour, _ := regexp.Compile(`^([A-Za-z]{1,16}) is ([\d]+) years (younger|older) than ([A-Za-z]{1,16})\n`)

	fmt.Fscan(in, &t)
	for t > 0 {
		line, err := in.ReadString('\n')
		if err != nil {
			break
		}
		if line == "\n" {
			continue
		}

		age := make(map[string]int, 3)
		relations := make([]assert, 0, 2)

		match := reQuestion.FindStringSubmatch(line)
		find := match[1]

		for i := 0; i < 3; {
			line, err := in.ReadString('\n')
			if err != nil {
				break
			}
			if line == "\n" {
				continue
			}

			i++

			// 1
			match := reOne.FindStringSubmatch(line)
			if len(match) == 3 {
				age[match[1]], _ = strconv.Atoi(match[2])
				continue
			}

			// 2
			match = reTwo.FindStringSubmatch(line)
			if len(match) == 3 {
				relations = append(relations, assert{match[1], 0, match[2]})
				continue
			}

			// 3, 4
			match = reThreeFour.FindStringSubmatch(line)
			if len(match) == 5 {
				odd, _ := strconv.Atoi(match[2])
				if match[3] == "younger" {
					odd *= -1
				}
				relations = append(relations, assert{match[1], odd, match[4]})
			}
		}

		if _, ok := age[find]; !ok {

			for i := 0; i < len(relations); {
				if v, ok := age[relations[i].whom]; ok {
					age[relations[i].who] = v + relations[i].old
				} else if v, ok := age[relations[i].who]; ok {
					age[relations[i].whom] = v - relations[i].old
				} else {
					// Если для текущей связи не может выявить возраст ...
					if i == len(relations) { // ... это последняя строка, тогда выходим из цикла
						break
					}
					// Меняем местами, т.к одна из двух строк, может быть решением
					relations[i], relations[i+1] = relations[i+1], relations[i]
					continue
				}
				i++
			}
		}

		fmt.Fprintln(out, age[find])

		t--
	}
}
