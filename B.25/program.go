package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func Do(in io.Reader, out io.Writer) {
	var nailCount int
	var nail uint16
	fmt.Fscan(in, &nailCount)
	nails := make([]uint16, nailCount, nailCount)
	for i := 0; i < nailCount; i++ {
		fmt.Fscan(in, &nail)
		nails[i] = nail
	}
	fmt.Fprintln(out, nails)
	sort.Slice(nails, func(i int, j int) bool { return nails[i] < nails[j] })
	fmt.Fprintln(out, nails)
	var result int
	result = int(nails[1] - nails[0])
	if nailCount == 2 {
		fmt.Fprintln(out, result)
		return
	}
	fmt.Fprintln(out, "1st")
	fmt.Fprintln(out, result)
	last := nails[len(nails)-1] - nails[len(nails)-2]
	result += int(last)
	//	fmt.Fprintln(out, result)
	for i := 2; i < len(nails)-2; {
		left := nails[i] - nails[i-1]
		right := nails[i+1] - nails[i]
		if left < right {
			result += int(left)
			i++
			fmt.Fprintln(out, left)
		} else {
			result += int(right)
			i += 2
			fmt.Fprintln(out, right)
		}
		fmt.Fprintln(out, result)
		fmt.Fprintln(out, "----------")

	}
	fmt.Fprintln(out, "lasr")
	fmt.Fprintln(out, last)
	fmt.Fprintln(out, result)

}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Do(in, out)
}
