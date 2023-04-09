package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
)

type Vertex struct {
	neig  []int
	color int
}

func NewVertex() Vertex {
	v := Vertex{}
	v.neig = make([]int, 1, 1)
	return v
}

type queue []int

func (q queue) Push(v int) queue {
	return append(q, v)
}

func (q queue) Pop() (queue, int, error) {
	l := len(q)
	if l == 0 {
		return q, 0, errors.New("empty")
	}
	return q[1:l], q[0], nil
}

func (q queue) Front() (int, error) {
	if len(q) == 0 {
		return 0, errors.New("empty")
	}
	return q[0], nil
}

func (q queue) Size() int {
	return len(q)
}

var vertex map[int]Vertex
var FLAG int
var result []int
var start int
var end int

func PrepareData(dataChan <-chan []int, signalChan chan<- int) {
	for v := range dataChan {
		if val, ok := vertex[v[0]]; ok {
			val.neig = append(val.neig, v[1])
			vertex[v[0]] = val
		} else {
			tmp := NewVertex()
			tmp.neig = append(tmp.neig, v[1])
			vertex[v[0]] = tmp
		}
	}
	signalChan <- 0
}

func Do(start int, end int, signalChan chan<- int) {
	q := queue{}
	q = q.Push(start)
	for q.Size() != 0 {
		var v int
		var e error
		q, v, e = q.Pop()
		if e == nil {
			if v == end {
				signalChan <- vertex[v].color
				return
			} else {
				currentVtx := vertex[v]
				for i := 1; i < len(currentVtx.neig); i++ {
					neigVtx := vertex[currentVtx.neig[i]]
					if neigVtx.color == 0 {
						neigVtx.color = currentVtx.color + 1
						q = q.Push(currentVtx.neig[i])
						vertex[currentVtx.neig[i]] = neigVtx
					}
				}
			}
		}
	}
	signalChan <- -1
}

func DoInput(in io.Reader, out io.Writer) {
	var n int
	FLAG = -1
	fmt.Fscan(in, &n)
	dataChan := make(chan []int)
	signalChan := make(chan int)
	vertex = make(map[int]Vertex)
	go PrepareData(dataChan, signalChan)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var v int
			fmt.Fscan(in, &v)
			if v != 0 {
				tmp := make([]int, 3, 3)
				tmp[0] = i + 1
				tmp[1] = j + 1
				tmp[2] = v
				dataChan <- tmp
			}
		}
	}
	close(dataChan)
	fmt.Fscan(in, &start, &end)
	<-signalChan
	go Do(start, end, signalChan)
	fmt.Fprintln(out, <-signalChan)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	DoInput(in, out)
}
