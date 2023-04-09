package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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

func isOperation(r rune) bool {
	switch r {
	case '+', '-', '*':
		return true
	}
	return false
}

func Do(in io.Reader, out io.Writer) {
	var str string
	s := stack{}
	var err error
	err = nil
	for err == nil {
		_, err = fmt.Fscan(in, &str)
		if err != nil {
			break
		}
		//		fmt.Fprintln(out, str)
		run := []rune(str)
		r := run[0]
		if r == ' ' {
			continue
		}
		switch isOperation(r) {
		case true:
			var op1, op2 int
			s, op1 = s.Pop()
			s, op2 = s.Pop()
			switch r {
			case '+':
				s = s.Push(op2 + op1)
			case '-':
				s = s.Push(op2 - op1)
			case '*':
				s = s.Push(op2 * op1)
			}
		case false:
			value, _ := strconv.Atoi(str)
			s = s.Push(value)
		}
		//fmt.Fprintln(out, s)
	}
	s, res := s.Pop()
	fmt.Fprintln(out, res)

}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Do(in, out)
}
