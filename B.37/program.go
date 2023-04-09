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
			tmp.color = -1
			tmp.neig = append(tmp.neig, v[1])
			vertex[v[0]] = tmp
		}
	}
	signalChan <- 0
}

func Do(start int, end int, signalChan chan<- int) int {
	q := queue{}
	vtx := vertex[start]
	vtx.color++
	vertex[start] = vtx
	q = q.Push(start)
	for q.Size() != 0 {
		var v int
		var e error
		q, v, e = q.Pop()
		if e == nil {
			if v == end {
				return vertex[v].color
			} else {
				currentVtx := vertex[v]
				for i := 1; i < len(currentVtx.neig); i++ {
					neigVtx := vertex[currentVtx.neig[i]]
					if neigVtx.color == -1 {
						neigVtx.color = currentVtx.color + 1
						q = q.Push(currentVtx.neig[i])
						vertex[currentVtx.neig[i]] = neigVtx
					}
				}
			}
		}
	}
	return -1
}

func getPath(end int) {
	for i := len(result) - 1; i >= 0; i-- {
		result[i] = end
		currentVtx := vertex[end]
		for j := 1; j < len(currentVtx.neig); j++ {
			if vertex[currentVtx.neig[j]].color == currentVtx.color-1 {
				end = currentVtx.neig[j]
			}
		}
	}
}

func DoInput(in io.Reader, out io.Writer) {
	var n int
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
	v := Do(start, end, signalChan)

	fmt.Fprintln(out, v)
	if v > 0 {
		result = make([]int, v+1, v+1)
		getPath(end)
		for _, v := range result {
			fmt.Fprintf(out, "%d ", v)
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	DoInput(in, out)
}
