package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"sort"
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
var count int
var colors []map[int]bool

func PrepareData(dataChan <-chan []int, signalChan chan<- int) {
	for v := range dataChan {
		vertex[v[0]].neig = append(vertex[v[0]].neig, v[1])
		vertex[v[1]].neig = append(vertex[v[1]].neig, v[0])

	}
	signalChan <- 0
}

func Do(vertexNum int) {
	if vertex[vertexNum].color != 0 {
		return
	}
	vertex[vertexNum].color = len(colors) - 1
	colors[len(colors)-1][vertexNum] = true
	for i := 1; i < len(vertex[vertexNum].neig); i++ {
		Do(vertex[vertexNum].neig[i])
	}
}

func DoInput(in io.Reader, out io.Writer) {
	var n int
	var m int
	fmt.Fscan(in, &n, &m)
	dataChan := make(chan []int)
	signalChan := make(chan int)
	vertex = make([]Vertex, n+1, n+1)
	colors = make([]map[int]bool, 1, 1)
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
	for i := 1; i < len(vertex); i++ {
		if vertex[i].color == 0 {
			colors = append(colors, make(map[int]bool))
			Do(i)
		}
	}
	fmt.Fprintln(out, len(colors)-1)
	for i := 1; i < len(colors); i++ {
		fmt.Fprintln(out, len(colors[i]))
		keys := make([]int, 0, len(colors[i]))
		for k := range colors[i] {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		for _, k := range keys {
			fmt.Fprintf(out, "%d ", k)
		}
		fmt.Fprintln(out)
	}

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	DoInput(in, out)
}
