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
	data := []byte{0, 1, 2, 3}
	var tmp []byte
	shab := byte(4)
	mask := byte(7)
	//if n < 27 {
	for i := uint64(3); i <= n; i++ {
		tmp = make([]byte, len(data))
		copy(tmp, data)
		for j := 0; j < len(tmp); j++ {
			tmp[j] = tmp[j] | shab
		}
		for j := 0; j < len(tmp); j++ {
			if tmp[j]&mask != mask {
				data = append(data, tmp[j])
				result++
			}
		}
		for i, _ := range data {
			data[i] = data[i] >> 1
		}
	}
	fmt.Fprintf(out, "%d\n", result)
	//}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Do(in, out)
}
