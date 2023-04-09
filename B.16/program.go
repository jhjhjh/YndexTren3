package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func Do(in io.Reader, out io.Writer) {
	var str string
	q := queue{}
	for {
		fmt.Fscan(in, &str)
		//		fmt.Fprintln(out, str)
		switch str {
		case "push":
			{
				var i int
				fmt.Fscan(in, &i)
				q = q.Push(i)
				fmt.Fprintln(out, "ok")
			}
		case "pop":
			if q.Size() > 0 {
				var i int
				q, i = q.Pop()
				fmt.Fprintln(out, i)
			} else {
				fmt.Fprintln(out, "error")
			}
		case "front":
			if q.Size() > 0 {
				fmt.Fprintln(out, q.Front())
			} else {
				fmt.Fprintln(out, "error")
			}
		case "size":
			fmt.Fprintln(out, q.Size())
		case "clear":
			q = q.Clear()
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
