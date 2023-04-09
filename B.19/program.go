package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type deque []int

func (d deque) PushBack(v int) deque {
	return append(d, v)
}

func (d deque) PushFront(v int) deque {
	return append([]int{v}, d...)
}

func (d deque) PopFront() (deque, int, error) {
	if len(d) == 0 {
		return d, 0, errors.New("empty")
	}
	l := len(d)
	return d[1:l], d[0], nil
}

func (d deque) PopBack() (deque, int, error) {
	if len(d) == 0 {
		return d, 0, errors.New("empty")
	}
	l := len(d)
	return d[0 : l-1], d[l-1], nil
}

func (d deque) Front() (int, error) {
	if len(d) == 0 {
		return 0, errors.New("empty")
	}
	return d[0], nil
}

func (d deque) Back() (int, error) {
	if len(d) == 0 {
		return 0, errors.New("empty")
	}
	return d[len(d)-1], nil
}

func (d deque) Size() int {
	return len(d)
}

func (d deque) Clear() deque {
	return deque{}
}

func (d deque) Insert(v int) deque {
	index := len(d)
	d = append(d, v)
	for index != 0 {
		tmp := (index - 1) >> 1
		if d[index] > d[tmp] {
			d[index], d[tmp] = d[tmp], d[index]
			index = tmp
		} else {
			break
		}
	}
	return d
}

func (d deque) print(out io.Writer) {
	fmt.Fprintln(out, d)
}

func (d deque) Extract() (deque, int, error) {
	if len(d) == 0 {
		return d, 0, errors.New("empty")
	}
	tmp := d[0]
	d[0] = d[len(d)-1]
	index := 0
	for index < len(d)-1 {
		tmpLeft := (index << 1) + 1
		if tmpLeft >= len(d) {
			break
		}
		tmpRight := tmpLeft + 1
		if tmpRight >= len(d) {
			if d[index] < d[tmpLeft] {
				d[index], d[tmpLeft] = d[tmpLeft], d[index]
				index = tmpLeft
			} else {
				break
			}
		} else {
			var tmpIdx int
			if d[tmpLeft] < d[tmpRight] {
				tmpIdx = tmpRight
			} else {
				tmpIdx = tmpLeft
			}
			if d[tmpIdx] > d[index] {
				d[tmpIdx], d[index] = d[index], d[tmpIdx]
				index = tmpIdx
			} else {
				break
			}

		}
	}
	return d[:len(d)-1], tmp, nil

}

func Do(in io.Reader, out io.Writer) {
	var str string
	d := deque{}
	var commandCount int
	fmt.Fscan(in, &commandCount)
	for i := 0; i < commandCount; i++ {
		fmt.Fscan(in, &str)
		switch str {
		case "0":
			var arg int
			fmt.Fscan(in, &arg)
			d = d.Insert(arg)
		case "1":
			var res int
			d, res, _ = d.Extract()
			fmt.Fprintln(out, res)
		}
		//d.print(out)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Do(in, out)
}
