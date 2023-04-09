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
	s := stack{}
	var vCount int
	fmt.Fscan(in, &vCount)
	vagons := stack{}
	result := stack{}
	tmp := make([]int, vCount, vCount)
	for i := 0; i < vCount; i++ {
		var mL int
		fmt.Fscan(in, &mL)
		tmp[i] = mL
	}
	for i := len(tmp) - 1; i >= 0; i-- {
		vagons = vagons.Push(tmp[i])
	}
	for vagons.Size() > 0 {
		var t int
		if s.Size() == 0 {
			vagons, t = vagons.Pop()
			s = s.Push(t)
		} else if vagons.Back() < s.Back() {
			vagons, t = vagons.Pop()
			s = s.Push(t)
		} else {
			s, t = s.Pop()
			result = result.Push(t)
		}
	}
	for s.Size() > 0 {
		var t int
		s, t = s.Pop()
		result = result.Push(t)
	}
	for i := 1; i < len(result); i++ {
		if result[i-1] > result[i] {
			fmt.Fprintln(out, "NO")
			return
		}
	}
	fmt.Fprintln(out, "YES")
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Do(in, out)
}
