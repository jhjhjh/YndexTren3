package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type queue []int

func (q queue) Push(v int) queue {
	return append(q, v)
}

func (q queue) Pop() (queue, int) {
	l := len(q)
	return q[1:l], q[0]
}

func (q queue) Front() int {
	return q[0]
}

func (q queue) Size() int {
	return len(q)
}

func (q queue) Clear() queue {
	return queue{}
}

func (q queue) print(out io.Writer) {
	fmt.Fprintln(out, q)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Do(in *bufio.Scanner, out io.Writer) {
	var str string
	first := queue{}
	second := queue{}
	in.Scan()
	str = in.Text()
	runes := []rune(str)
	for _, v := range runes {
		if v == ' ' {
			continue
		}
		num, _ := strconv.Atoi(string(v))
		first = first.Push(num)
	}
	in.Scan()
	str = in.Text()
	runes = []rune(str)
	for _, v := range runes {
		if v == ' ' {
			continue
		}
		num, _ := strconv.Atoi(string(v))
		second = second.Push(num)
	}
	for i := 1; i < 1000001; i++ {
		var fst, scd int
		first, fst = first.Pop()
		second, scd = second.Pop()
		razn := fst - scd
		if abs(razn) == 9 {
			if razn > 0 {
				second = second.Push(fst)
				second = second.Push(scd)
			} else {
				first = first.Push(fst)
				first = first.Push(scd)
			}
		} else {
			if razn < 0 {
				second = second.Push(fst)
				second = second.Push(scd)
			} else {
				first = first.Push(fst)
				first = first.Push(scd)
			}
		}
		if first.Size() == 0 {
			fmt.Fprint(out, "second ")
			fmt.Fprintln(out, i)
			return
		}
		if second.Size() == 0 {
			fmt.Fprint(out, "first ")
			fmt.Fprintln(out, i)
			return
		}
	}
	fmt.Fprintln(out, "botva")

}

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Do(in, out)
}
