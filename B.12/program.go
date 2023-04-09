package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type stack []rune

func (s stack) Push(v rune) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, rune) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Back() rune {
	return s[len(s)-1]
}

func (s stack) Size() int {
	return len(s)
}

func (s stack) Clear() stack {
	return stack{}
}

type parenth struct {
	fl map[rune]rune
}

func Do(in io.Reader, out io.Writer) {
	var str string
	var open = map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	s := stack{}
	fmt.Fscan(in, &str)
	runes := []rune(str)
	flag := true
	for i := 0; i < len(runes); i++ {
		switch runes[i] {
		case '(', '{', '[':
			s = s.Push(runes[i])
		default:
			if s.Size() == 0 {
				flag = false
			} else if s.Back() == open[runes[i]] {
				s, _ = s.Pop()
			} else if s.Back() == '(' || s.Back() == '{' || s.Back() == '[' {
				flag = false
			} else {
				s = s.Push(runes[i])
			}

		}
		if !flag {
			break
		}
		//	fmt.Fprintln(out, flag)
		//	fmt.Fprintln(out, s)
	}
	//	fmt.Fprintln(out, s)
	if flag && s.Size() == 0 {
		fmt.Fprint(out, "yes")
	} else {
		fmt.Fprint(out, "no")
	}

}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Do(in, out)
}
