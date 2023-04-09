package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Do(in io.Reader, out io.Writer) {
	var n int32
	fmt.Fscan(in, &n)
	data := make([][]int32, 0, 0)
	data = append(data, make([]int32, 0, 0))
	data[0] = append(data[0], n)
	step := 0
	Stop := false
	index := 0
	for !Stop && n > 1 {
		step++
		data = append(data, make([]int32, 0, 0))
		for i := 0; i < len(data[step-1]) && !Stop; i++ {
			tmp := data[step-1][i]
			if tmp%3 == 0 {
				data[step] = append(data[step], tmp/3)
				if data[step][len(data[step])-1] == 1 {
					Stop = true
					index = len(data[step]) - 1
				}
			}
			if tmp%2 == 0 {
				data[step] = append(data[step], tmp/2)
				if data[step][len(data[step])-1] == 1 {
					Stop = true
					index = len(data[step]) - 1
				}
			}
			data[step] = append(data[step], tmp-1)
			if data[step][len(data[step])-1] == 1 {
				Stop = true
				index = len(data[step]) - 1
			}

		}

	}
	fmt.Fprintf(out, "%d\n", step)
	//fmt.Fprintln(out, data)
	for i := step; i > 0; i-- {
		current := data[i][index]
		fmt.Fprintf(out, "%d ", current)
		for j := 0; j < len(data[i-1])-1; j++ {
			if current*3 == data[i-1][j] {
				index = j
				break
			}
			if current*2 == data[i-1][j] {
				index = j
				break
			}
			if current+1 == data[i-1][j] {
				index = j
			}
		}
	}
	fmt.Fprintln(out, data[0][0])
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Do(in, out)
}
