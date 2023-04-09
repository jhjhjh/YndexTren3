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

func PrepareData(dataChan <-chan []int, signalChan chan<- int) {
	for v := range dataChan {
		vertex[v[0]].neig = append(vertex[v[0]].neig, v[1])
		vertex[v[1]].neig = append(vertex[v[1]].neig, v[0])

	}
	signalChan <- 0
}

func Do(vertexNum int, color int) {
	if vertex[vertexNum].color == 0 {
		vertex[vertexNum].color = 3 - color
		for i := 1; i < len(vertex[vertexNum].neig) && YES; i++ {
			Do(vertex[vertexNum].neig[i], vertex[vertexNum].color)
		}
	} else if vertex[vertexNum].color == color {
		YES = false
	}
}

func DoInput(in io.Reader, out io.Writer) {
	var n int
	var m int
	YES = true
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
	color := 1
	for i := 1; i < len(vertex); i++ {
		if vertex[i].color == 0 {
			Do(i, color)
		}
	}
	if YES {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	DoInput(in, out)
}
