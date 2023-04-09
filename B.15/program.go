package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type stack []int32

func (s stack) Push(v int32) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int32) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Back() int32 {
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
	var cityC int
	fmt.Fscan(in, &cityC)
	cities := make([]int32, cityC, cityC)
	result := make([]int32, cityC, cityC)
	for i := int32(0); i < int32(cityC); i++ {
		var mL int32
		fmt.Fscan(in, &mL)
		cities[i] = mL

	}
	s = s.Push(int32(cityC) - 1)
	result[len(result)-1] = -1
	for i := int32(cityC) - 2; i >= 0; i-- {
		if cities[i] > cities[s.Back()] {
			result[i] = s.Back()
			s = s.Push(i)
		} else {
			for cities[s.Back()] >= cities[i] {
				s, _ = s.Pop()
				if s.Size() == 0 {
					break
				}
			}
			if s.Size() == 0 {
				result[i] = -1
			} else {
				result[i] = s.Back()
			}
			s = s.Push(i)
		}
	}
	for _, v := range result {
		fmt.Fprint(out, v)
		fmt.Fprint(out, " ")
	}

}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Do(in, out)
}
