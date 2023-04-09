package main

import (
	"bufio"
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

func (d deque) PopFront() (deque, int) {
	l := len(d)
	return d[1:l], d[0]
}

func (d deque) PopBack() (deque, int) {
	l := len(d)
	return d[0 : l-1], d[l-1]
}

func (d deque) Front() int {
	return d[0]
}

func (d deque) Back() int {
	return d[len(d)-1]
}

func (d deque) Size() int {
	return len(d)
}

func (d deque) Clear() deque {
	return deque{}
}

func Do(in io.Reader, out io.Writer) {
	var str string
	d := deque{}
	for {
		fmt.Fscan(in, &str)
		//		fmt.Fprintln(out, str)
		switch str {
		case "push_front":
			{
				var i int
				fmt.Fscan(in, &i)
				d = d.PushFront(i)
				fmt.Fprintln(out, "ok")
			}
		case "push_back":
			{
				var i int
				fmt.Fscan(in, &i)
				d = d.PushBack(i)
				fmt.Fprintln(out, "ok")
			}
		case "pop_front":
			if d.Size() > 0 {
				var i int
				d, i = d.PopFront()
				fmt.Fprintln(out, i)
			} else {
				fmt.Fprintln(out, "error")
			}
		case "pop_back":
			if d.Size() > 0 {
				var i int
				d, i = d.PopBack()
				fmt.Fprintln(out, i)
			} else {
				fmt.Fprintln(out, "error")
			}
		case "front":
			if d.Size() > 0 {
				fmt.Fprintln(out, d.Front())
			} else {
				fmt.Fprintln(out, "error")
			}
		case "back":
			if d.Size() > 0 {
				fmt.Fprintln(out, d.Back())
			} else {
				fmt.Fprintln(out, "error")
			}
		case "size":
			fmt.Fprintln(out, d.Size())
		case "clear":
			d = d.Clear()
			fmt.Fprintln(out, "ok")

		}
		if str == "exit" {
			fmt.Fprint(out, "bye")
			break
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Do(in, out)
}
