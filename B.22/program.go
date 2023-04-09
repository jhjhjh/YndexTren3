package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type data []int

func (d data) getEl(a int) int {
	if a <= 0 {
		return 0
	}
	return d[a]
}

func Do(in io.Reader, out io.Writer) {
	var n int
	var k int
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &k)
	data := make(data, 3, 3)
	data[1] = 1
	data[2] = 1
	//fmt.Fprintln(out, data.getEl(1))
	for i := 3; i <= n; i++ {
		var tmp int
		for j := 1; j <= k; j++ {
			tmp += data.getEl(i - j)
			//fmt.Fprintf(out,"" tmp)
		}
		//fmt.Fprintln(out, tmp)
		data = append(data, tmp)
	}
	fmt.Fprintln(out, data[n])
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Do(in, out)
}
