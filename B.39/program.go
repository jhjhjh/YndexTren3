package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
)

var field [][][]int

var N int

var startPoint [3]int
var result int

func calcField() {
	q := queue{}
	q = q.Push(startPoint)
	for q.Size() > 0 {
		var currCell [3]int
		q, currCell, _ = q.Pop()
		if currCell[0] == 0 {
			result = getCell(currCell)
			return
		}
		genCellChan := make(chan [3]int)
		go func(out chan<- [3]int, currentCell [3]int) {
			v := [6][3]int{{-1, 0, 0}, {1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, -1}, {0, 0, 1}}
			for _, tt := range v {
				tt[0] += currentCell[0]
				tt[1] += currentCell[1]
				tt[2] += currentCell[2]
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

func getCell(cell [3]int) int {
	if cell[0] >= 0 && cell[0] < len(field) && cell[1] >= 0 && cell[1] < len(field[0]) && cell[2] >= 0 && cell[2] < len(field[0][0]) {
		return field[cell[0]][cell[1]][cell[2]]
	}
	return -2
}

func setCell(cell [3]int, c int) {
	if cell[0] >= 0 && cell[0] < len(field) && cell[1] >= 0 && cell[1] < len(field[0]) && cell[2] >= 0 && cell[2] < len(field[0][0]) {
		field[cell[0]][cell[1]][cell[2]] = c
	}
}

type queue [][3]int

func (q queue) Push(v [3]int) queue {
	return append(q, v)
}

func (q queue) Pop() (queue, [3]int, error) {
	l := len(q)
	if l == 0 {
		return q, [3]int{}, errors.New("empty")
	}
	return q[1:l], q[0], nil
}

func (q queue) Size() int {
	return len(q)
}

func prepareField(in <-chan data, signalChan chan<- int) {
	for i, _ := range field {
		field[i] = make([][]int, N, N)
		for j, _ := range field[i] {
			field[i][j] = make([]int, N, N)
		}
	}
	for data := range in {
		for i, v := range data.data {
			switch v {
			case 35:
				setCell([3]int{data.v1, data.v2, i}, -2)
			case 46:
				setCell([3]int{data.v1, data.v2, i}, -1)
			case 83:
				setCell([3]int{data.v1, data.v2, i}, 0)
				startPoint = [3]int{data.v1, data.v2, i}
			}
		}
	}
	signalChan <- 0
}

type data struct {
	v1   int
	v2   int
	data []byte
}

func DoInput(in io.Reader, out io.Writer) {
	fmt.Fscan(in, &N)
	field = make([][][]int, N, N)
	dataChan := make(chan data)
	signalChan := make(chan int)
	go prepareField(dataChan, signalChan)
	for i := 0; i < N; i++ {
		var inputString string
		fmt.Fscan(in, inputString)
		for j := 0; j < N; j++ {
			fmt.Fscan(in, &inputString)
			dt := data{v1: i, v2: j, data: []byte(inputString)}
			dataChan <- dt
		}
	}
	close(dataChan)
	<-signalChan
	calcField()
	fmt.Fprintln(out, result)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	DoInput(in, out)
}
