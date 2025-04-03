package main

import (
	"bufio"
	"fmt"
	"os"
)

type exchange struct {
	n, m float32
}

/*
0 - Курс обмена рублей на доллары.
1 - Курс обмена рублей на евро.
2 - Курс обмена долларов на рубли.
3 - Курс обмена долларов на евро.
4 - Курс обмена евро на рубли.
5 - Курс обмена евро на доллары.
*/
type bank struct {
	exchange []exchange
}

func max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		t int // кол-во тестов
	)

	// Возможные обмены, кроме прямого (RUB->USD)
	possibleExchage := [][]int{
		{1, 5},    // RUB->EUR->USD
		{1, 4, 0}, // RUB->EUR->RUB->USD
		{0, 2, 0}, // RUB->USD->RUB->USD
		{0, 3, 5}, // RUB->USD->EUR->USD
	}

	// Возможные пути обхода банков
	possibleBanksRouteExchange := [][]int{
		{0, 1, 2},
		{0, 2, 1},
		{1, 0, 2},
		{1, 2, 0},
		{2, 1, 0},
		{2, 0, 1},
	}

	fmt.Fscan(in, &t)
	for t > 0 {
		var bestExchange float32 // лучший курс обмена

		banks := make([]bank, 3)
		for b := 0; b < 3; b++ {
			banks[b].exchange = make([]exchange, 6)
			// rub -> usd
			fmt.Fscan(in, &banks[b].exchange[0].n, &banks[b].exchange[0].m)
			bestExchange = 1.0 * banks[b].exchange[0].m / banks[b].exchange[0].n // best RUB->USD

			for e := 1; e < 6; e++ {
				fmt.Fscan(in, &banks[b].exchange[e].n, &banks[b].exchange[e].m)
			}
		}

		for _, pE := range possibleExchage {
			// обмен с 1 рублём для всех возможных банков possibleBanksRouteExchange
			start := []float32{1.0, 1.0, 1.0, 1.0, 1.0, 1.0}

			for ie, e := range pE {
				for ib, b := range possibleBanksRouteExchange {
					start[ib] = start[ib] * banks[b[ie]].exchange[e].m / banks[b[ie]].exchange[e].n
				}
			}

			for _, v := range start {
				bestExchange = max(bestExchange, v)
			}
		}
		fmt.Fprintln(out, bestExchange)

		t--
	}
}
