package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type car struct {
	number   int
	start    int
	end      int
	capacity int
}

type arrival struct {
	index int
	time  int
	car   int
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		t int // кол-во наборов
		n int // кол-во заказов
	)

	fmt.Fscan(in, &t)
	//fmt.Fprintln(out, "T: ", t)

	for t > 0 {
		fmt.Fscan(in, &n)
		//fmt.Fprintln(out, "N: ", n)

		arrivals := make([]arrival, n, n) // момент времени прибытия заказа в пункт
		for i := 0; i < n; i++ {
			arrivals[i] = arrival{i, 0, -1}
			fmt.Fscan(in, &arrivals[i].time)
		}
		//fmt.Fprintln(out, "Arrival: ", arrival)
		sort.Slice(arrivals, func(i, j int) bool {
			return arrivals[i].time < arrivals[j].time
		})
		//fmt.Fprintln(out, "Sorted Arrival: ", arrival)
		//fmt.Fprintln(out, "Arrival Map: ", arrivalMap)

		var m int // кол-во грузовых машин
		fmt.Fscan(in, &m)
		//fmt.Fprintln(out, "M: ", m)

		cars := make([]car, m, m)
		for j := 0; j < m; j++ {
			cars[j] = car{j + 1, 0, 0, 0}
			fmt.Fscan(in, &cars[j].start, &cars[j].end, &cars[j].capacity)
		}
		//fmt.Fprintln(out, "Cars: ", cars)
		// Сортируем машины по времени прибытия
		sort.Slice(cars, func(i, j int) bool {
			// если время прибытие равное, то сортируем по номеру машины
			if cars[i].start == cars[j].start {
				return cars[i].number < cars[j].number
			}
			return cars[i].start < cars[j].start
		})
		//fmt.Fprintln(out, "Sort Cars: ", cars)

		indexCar := 0
		for i := 0; i < n; {
			// закончились машины
			if indexCar >= m {
				break
			}
			// если заказ пришёл раньше, чем пришла машина, то переходим к новому заказу
			if arrivals[i].time < cars[indexCar].start {
				i++
				continue
			}

			// в машине нет места или текущая машина уже уехала, переходим к следующей машине
			if cars[indexCar].capacity == 0 || arrivals[i].time > cars[indexCar].end {
				indexCar++
				continue
			}
			//fmt.Fprintln(out, "car:", cars[indexCar].number, "index:", i, "Time a:", arrival[i], "Map:", arrivalMap[arrival[i]])

			// Сохраняем номер машины в соответствующей ячейки заказа
			arrivals[i].car = cars[indexCar].number
			cars[indexCar].capacity--
			i++
		}

		// Сортируем по индексу заказазов, как они были указаны в наборе
		sort.Slice(arrivals, func(i, j int) bool {
			return arrivals[i].index < arrivals[j].index
		})
		for _, v := range arrivals {
			fmt.Fprint(out, v.car, " ")
		}
		fmt.Fprintln(out, "")
		t--

		// fmt.Fprintln(out, "")
	}
}
