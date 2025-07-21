package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type ourStruct struct {
	Dir     string      `json:"dir"`
	Files   []string    `json:"files,omitempty"`
	Folders []ourStruct `json:"folders"`
}

func search(data *ourStruct, isInfected bool, virus *int) {
	// fmt.Println(data.Dir, isInfected, data.Files, data.Folders)

	// Если родительская папка была инфецирована, то все файлы заражены
	if isInfected {
		// прибавляем кол-во зараженных файлов
		*virus += len(data.Files)
	} else {
		for _, filename := range data.Files {
			// если в название файла есть суфикс .hack, то все файлы заражены
			if !strings.HasSuffix(filename, ".hack") {
				continue
			}
			isInfected = true
			*virus += len(data.Files)
			break
		}
	}

	for _, folder := range data.Folders {
		search(&folder, isInfected, virus)
	}
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		t int // кол-во наборов
		n int // кол-во строк
	)
	// start := time.Now()
	fmt.Fscan(in, &t)
	for t > 0 {
		fmt.Fscan(in, &n)

		var jsonString bytes.Buffer
		for i := 0; i <= n; i++ {
			line, err := in.ReadString('\n')
			if line == "\n" {
				continue
			} else if err == io.EOF {
				break
			}
			jsonString.WriteString(line)
		}

		var data ourStruct
		err := json.Unmarshal(jsonString.Bytes(), &data)
		if err != nil {
			panic(err)
		}

		virus := 0
		search(&data, false, &virus)
		fmt.Fprintln(out, virus)

		t--
	}
	// fmt.Fprintln(out, "Time:", time.Since(start))
}
