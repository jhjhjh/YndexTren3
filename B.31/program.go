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
	color int8
}

var vertex []Vertex
var count int

func PrepareData(dataChan <-chan []int, signalChan chan<- int) {
	for v := range dataChan {
		//fmt.Println(v)
		if vertex[v[0]].neig == nil {
			vertex[v[0]].neig = make([]int, 1, 1)
		}
		vertex[v[0]].neig = append(vertex[v[0]].neig, v[1])
		if vertex[v[1]].neig == nil {
			vertex[v[1]].neig = make([]int, 1, 1)
		}
		vertex[v[1]].neig = append(vertex[v[1]].neig, v[0])

	}
	//fmt.Println(vertex)
	signalChan <- 0
}

func Do(vertexNum int) {
	//	fmt.Println()
	//	fmt.Println(vertex)
	if vertex[vertexNum].color != 0 {
		return
	}
	vertex[vertexNum].color = 1
	count++
	for i := 1; i < len(vertex[vertexNum].neig); i++ {
		//		fmt.Println(vertex[vertexNum].neig)
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
	go PrepareData(dataChan, signalChan)
	for i := 0; i < m; i++ {
		v := make([]int, 2, 2)
		fmt.Fscan(in, &v[0], &v[1])
		//	t.Fprintln(out, v)
		dataChan <- v
	}
	close(dataChan)
	<-signalChan
	Do(1)
	fmt.Fprintln(out, count)
	for i, _ := range vertex {
		if vertex[i].color == 1 {
			fmt.Fprintf(out, "%d ", i)
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
