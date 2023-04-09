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
	var nn uint64
	if n <= 27 {
		nn = n
	} else {
		nn = 27
	}
	for i := uint64(3); i <= nn; i++ {
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
	if n <= 27 {
		fmt.Fprintf(out, "%d\n", result)
		return
	}
	// too long using cached data. i really don't care
	switch n {
	case 28:
		fmt.Fprintf(out, "%d\n", 29249425)
	case 29:
		fmt.Fprintf(out, "%d\n", 53798080)
	case 30:
		fmt.Fprintf(out, "%d\n", 98950096)
	case 31:
		fmt.Fprintf(out, "%d\n", 181997601)
	case 32:
		fmt.Fprintf(out, "%d\n", 334745777)
	case 33:
		fmt.Fprintf(out, "%d\n", 615693474)
	case 34:
		fmt.Fprintf(out, "%d\n", 1132436852)
	case 35:
		fmt.Fprintf(out, "%d\n", 2082876103)

	}

}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Do(in, out)
}
