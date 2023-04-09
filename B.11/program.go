package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Back() int {
	return s[len(s)-1]
}

func (s stack) Size() int {
	return len(s)
}

func (s stack) Clear() stack {
	return stack{}
}

func Do(in io.Reader, out io.Writer) {
	var str string
	s := stack{}
	for {
		fmt.Fscan(in, &str)
		//		fmt.Fprintln(out, str)
		switch str {
		case "push":
			{
				var i int
				fmt.Fscan(in, &i)
				s = s.Push(i)
				fmt.Fprintln(out, "ok")
			}
		case "pop":
			if s.Size() > 0 {
				var i int
				s, i = s.Pop()
				fmt.Fprintln(out, i)
			} else {
				fmt.Fprintln(out, "error")
			}
		case "back":
			if s.Size() > 0 {
				fmt.Fprintln(out, s.Back())
			} else {
				fmt.Fprintln(out, "error")
			}
		case "size":
			fmt.Fprintln(out, s.Size())
		case "clear":
			s = s.Clear()
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
