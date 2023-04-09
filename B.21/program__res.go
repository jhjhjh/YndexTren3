package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func Do(in io.Reader, out io.Writer) {
	var n uint64
	fmt.Fscan(in, &n)
	//previos := uint64(4)
	result := 4
	if n == 1 {
		fmt.Fprintln(out, 2)
		return
	} else if n == 2 {
		fmt.Fprintln(out, 4)
		return
	}
	//current := uint64(4)
	//for i := uint64(3); i <= n; i++ {
	//	fmt.Fprintf(out, "i = %d\n", i)
	//	current = current << 1
	for j := uint64(4); j < uint64(math.Pow(float64(2), float64(n))); j++ {
		fmt.Fprintf(out, "\tj = %d\t%b\n", j, j)
		tmp := j
		flag := true
		for k := uint64(0); k < j; k++ {
			if tmp&7 == 7 {
				fmt.Fprintf(out, "1\t\t\t\t%b\n", tmp)
				flag = false
				break
			} else {
				fmt.Fprintf(out, "2\t\t\t\t%b\n", tmp)
			}
			tmp = tmp >> 1
			if tmp == 0 {
				break
			}

		}
		if flag {
			result++
		}
		fmt.Fprintf(out, "\t\t\tmiddle: %d\n", result)

	}
	//	previos = current
	//	fmt.Fprintf(out, "middle: %d\n", result)
	//}
	fmt.Fprintf(out, "%d\n", result)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Do(in, out)
}
