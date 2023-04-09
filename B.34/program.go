package main

import (
	"bufio"
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

var vertex []Vertex
var YES bool
var result []int

func PrepareData(dataChan <-chan []int, signalChan chan<- int) {
	for v := range dataChan {
		vertex[v[0]].neig = append(vertex[v[0]].neig, v[1])
		//vertex[v[1]].neig = append(vertex[v[1]].neig, v[0])

	}
	signalChan <- 0
}

func Do(vertexNum int) {
	switch vertex[vertexNum].color {
	case 0:
		vertex[vertexNum].color = 1
		for i := 1; i < len(vertex[vertexNum].neig) && YES; i++ {
			Do(vertex[vertexNum].neig[i])
		}
	case 1:
		YES = false
		return
	case 2:
		return
	}
	vertex[vertexNum].color = 2
	result = append(result, vertexNum)
}

func DoInput(in io.Reader, out io.Writer) {
	var n int
	var m int
	YES = true
	result = make([]int, 0, 0)
	fmt.Fscan(in, &n, &m)
	dataChan := make(chan []int)
	signalChan := make(chan int)
	vertex = make([]Vertex, n+1, n+1)
	go PrepareData(dataChan, signalChan)
	for i := 0; i <= n; i++ {
		vertex[i] = NewVertex()
	}
	for i := 0; i < m; i++ {
		v := make([]int, 2, 2)
		fmt.Fscan(in, &v[0], &v[1])
		dataChan <- v
	}
	close(dataChan)
	<-signalChan
	//fmt.Fprintln(out, vertex)
	for i := 1; i < len(vertex); i++ {
		if vertex[i].color == 0 {
			Do(i)
		}
	}
	if YES {
		for i := len(result) - 1; i >= 0; i-- {
			fmt.Fprintf(out, "%d ", result[i])
		}
	} else {
		fmt.Fprintln(out, "-1")
	}
	//fmt.Fprintln(out, result)

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	DoInput(in, out)
}
