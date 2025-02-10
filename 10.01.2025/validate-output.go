package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

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

		// Заполняем слайс прочитанными элементами из второй строки набора
		nums := make([]int, n, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &nums[i])
		}

		// Сортируем полученый слайс
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})

		// Формируем строку из полученного слайса, что бы потом её сравнить со строкой из файла
		sliceStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(nums)), " "), "[]") + "\n"

		// Считываем последнюю строку из набора
		var resultStr string
		for {
			resultStr, _ = in.ReadString('\n')
			if resultStr == "\n" {
				continue
			}
			break
		}

		// Сравниваем посленюю строку из набора со стракой отсортированного слайса
		if resultStr == sliceStr {
			fmt.Fprintln(out, "yes")
		} else {
			fmt.Fprintln(out, "no")
		}

		t--
	}
}
