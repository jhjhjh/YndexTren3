package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
)

var field [][]int

func DoInput(in io.Reader, out io.Writer) {
	var N, M int
	fmt.Fscan(in, &N, &M)
	field = make([][]int, N, N)
	for i := 0; i < N; i++ {
		field[i] = make([]int, M, M)
	}
	for i := 0; i < N; i++ {
		var tmp int
		for j := 0; j < M; j++ {
			fmt.Fscan(in, &tmp)
			field[i][j] = tmp
		}
	}
	for j := 1; j < M; j++ {
		field[0][j] += field[0][j-1]
	}
	for i := 1; i < N; i++ {
		field[i][0] += field[i-1][0]
		for j := 1; j < M; j++ {
			l := field[i][j] + field[i-1][j]
			u := field[i][j] + field[i][j-1]
			if l < u {
				field[i][j] = l
			} else {
				field[i][j] = u
			}
		}
	}
	fmt.Fprintln(out, field[N-1][M-1])
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	DoInput(in, out)
}
