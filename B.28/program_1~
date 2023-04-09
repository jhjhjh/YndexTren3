package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
)

var field [][]int

var N int
var M int

func NewField() {
	field = make([][]int, N)
	for i, _ := range field {
		field[i] = make([]int, M)
		for j, _ := range field[i] {
			field[i][j] = 0
		}
	}
}

func calcField(startCell [2]int) {
	q := queue{}
	q = q.Push(startCell)
	for q.Size() > 0 {
		var currCell [2]int
		q, currCell, _ = q.Pop()
		genCellChan := make(chan [2]int)
		go func(out chan<- [2]int, currentCell [2]int) {
			v := [2][2]int{{2, 1}, {1, 2}}
			for _, tt := range v {
				tt[0] += currentCell[0]
				tt[1] += currentCell[1]
				out <- tt
			}
			close(out)
		}(genCellChan, currCell)
		for gCell := range genCellChan {
			if getCell(gCell) >= -1 {
				m2 := getCell(gCell)
				setCell(gCell, m2+1)
				q = q.Push(gCell)
			}
		}
	}
}

func getCell(cell [2]int) int {
	if cell[0] < len(field) && cell[1] < len(field[0]) {
		return field[cell[0]][cell[1]]
	}
	return -2
}

func setCell(cell [2]int, c int) {
	if cell[0] < len(field) && cell[1] < len(field[0]) {
		field[cell[0]][cell[1]] = c
	}
}

type queue [][2]int

func (q queue) Push(v [2]int) queue {
	return append(q, v)
}

func (q queue) Pop() (queue, [2]int, error) {
	l := len(q)
	if l == 0 {
		return q, [2]int{}, errors.New("empty")
	}
	return q[1:l], q[0], nil
}

func (q queue) Size() int {
	return len(q)
}

func DoInput(in io.Reader, out io.Writer) {
	fmt.Fscan(in, &N, &M)
	NewField()
	calcField([2]int{0, 0})
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if field[i][j] > 0 {
				fmt.Fprintf(out, "%d;", field[i][j])
			} else {
				fmt.Fprintf(out, " ;")

			}
		}
		fmt.Fprintln(out)
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
