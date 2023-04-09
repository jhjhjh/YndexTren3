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
			field[i][j] = -1
		}
	}
}

func calcField(startCell [2]int) {
	field[startCell[0]-1][startCell[1]-1] = 0
	q := queue{}
	q = q.Push(startCell)
	for q.Size() > 0 {
		var currCell [2]int
		q, currCell, _ = q.Pop()
		genCellChan := make(chan [2]int)
		go func(out chan<- [2]int, currentCell [2]int) {
			v := [8][2]int{{-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {1, -2}, {2, -1}}
			for _, tt := range v {
				tt[0] += currentCell[0]
				tt[1] += currentCell[1]
				out <- tt
			}
			close(out)
		}(genCellChan, currCell)
		for gCell := range genCellChan {
			if getCell(gCell) == -1 {
				setCell(gCell, getCell(currCell)+1)
				q = q.Push(gCell)
			}
		}
	}
}

func getCell(cell [2]int) int {
	if cell[0]-1 >= 0 && cell[0]-1 < len(field) && cell[1]-1 >= 0 && cell[1]-1 < len(field[0]) {
		return field[cell[0]-1][cell[1]-1]
	}
	return -2
}

func setCell(cell [2]int, c int) {
	if cell[0]-1 >= 0 && cell[0]-1 < len(field) && cell[1]-1 >= 0 && cell[1]-1 < len(field[0]) {
		field[cell[0]-1][cell[1]-1] = c
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
	var s, t, q int
	fmt.Fscan(in, &N, &M, &s, &t, &q)
	NewField()
	calcField([2]int{s, t})
	var result int
	for i := 0; i < q; i++ {
		var v [2]int
		fmt.Fscan(in, &v[0], &v[1])
		if result != -1 {
			tmp := getCell(v)
			if tmp == -1 {
				result = -1
			} else {
				result += tmp
			}
		}
	}
	fmt.Fprintln(out, result)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	DoInput(in, out)
}
