package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Do(in io.Reader, out io.Writer) {
	var n uint64
	fmt.Fscan(in, &n)
	result := int64(4)
	if n == 1 {
		result = 2
	} else if n == 2 {
		result = 4
	}
	data := []byte{2, 3}
	shab := byte(4)
	mask := byte(7)
	for i := uint64(3); i <= n; i++ {
		//middle := len(data)
		//shab = shab << 1
		data = append(data, data...)
		for j := 0; j < len(data); j++ {
			//data[j] = data[j] >> 1
			data[j] = data[j] | shab
		}
		fmt.Fprintln(out, data)
		tmp := make([]byte, 0, len(data))
		//tmp = append(data[:middle])
		for j := 0; j < len(data); j++ {
			//fmt.Fprintf(out, " %b %b %b \n", data[j], mask, data[j]&mask)
			if data[j]&mask != mask {
				tmp = append(tmp, data[j])
				result++
			}
		}
		data = tmp
		for i, _ := range data {
			data[i] = data[i] >> 1
		}
		//mask = mask << 1
	}
	fmt.Fprintln(out, data)
	fmt.Fprintf(out, "%d\n", result)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Do(in, out)
}
