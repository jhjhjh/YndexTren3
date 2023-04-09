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

var vertex map[int]Vertex
var FLAG int
var result []int

func PrepareData(dataChan <-chan []int, signalChan chan<- int) {
	for v := range dataChan {
		if v[2] != 0 {
			if val, ok := vertex[v[0]]; ok {
				val.neig = append(val.neig, v[1])
				vertex[v[0]] = val
			} else {
				tmp := NewVertex()
				tmp.neig = append(tmp.neig, v[1])
				vertex[v[0]] = tmp
			}
		}
	}
	signalChan <- 0
}

func Do(vertexNum int, parent int) {
	if val, ok := vertex[vertexNum]; ok && FLAG < 0 {
		switch val.color {
		case 0:
			//fmt.Printf("vertexNum = %d  parent =  %d \n", vertexNum, parent)
			val.color = 1
			vertex[vertexNum] = val
			for i := 1; i < len(val.neig); i++ {
				if val.neig[i] != parent {
					Do(val.neig[i], vertexNum)
				}
			}
			val.color = 2
			vertex[vertexNum] = val
			if FLAG > 0 {
				result = append(result, vertexNum)
			}
			if FLAG == vertexNum {
				FLAG = 0
			}
		case 1:
			FLAG = vertexNum
			//fmt.Printf("vertexNum = %d, FLAG = %d", vertexNum, FLAG)
		}
	}
}

func DoInput(in io.Reader, out io.Writer) {
	var n int
	FLAG = -1
	result = make([]int, 0, 0)
	fmt.Fscan(in, &n)
	dataChan := make(chan []int)
	signalChan := make(chan int)
	vertex = make(map[int]Vertex)
	go PrepareData(dataChan, signalChan)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tmp := make([]int, 3, 3)
			tmp[0] = i + 1
			tmp[1] = j + 1
			fmt.Fscan(in, &tmp[2])
			dataChan <- tmp
		}
	}
	close(dataChan)
	<-signalChan
	//	fmt.Fprintln(out, vertex)
	for i := 1; i < len(vertex); i++ {
		if vertex[i].color == 0 {
			Do(i, -1)
		}
	}
	if FLAG == -1 {
		fmt.Fprintln(out, "NO")
	} else {
		fmt.Fprintln(out, "YES")
		fmt.Fprintln(out, len(result))
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
