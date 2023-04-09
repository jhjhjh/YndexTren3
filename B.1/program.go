package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var input string

	stat := make(map[byte]int)
	var err error
	for err == nil {
		input, err = in.ReadString('\n')
		input = strings.ReplaceAll(input, " ", "")
		input = strings.ReplaceAll(input, "\n", "")
		byteStr := []byte(input)
		for _, v := range byteStr {
			stat[v] = stat[v] + 1
		}
	}
	keys := make([]byte, 0, 26)
	var maxVal int
	for k, v := range stat {
		keys = append(keys, k)
		if v > maxVal {
			maxVal = v
		}
	}

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for i := 0; i < maxVal; i++ {
		for j := 0; j < len(keys); j++ {
			if stat[keys[j]]-maxVal+i >= 0 {
				fmt.Fprint(out, "#")
			} else {
				fmt.Fprint(out, " ")
			}
		}
		fmt.Fprintln(out)
	}
	for _, v := range keys {
		fmt.Fprint(out, string(v))
	}
	fmt.Fprintln(out)
}
