package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
)

type queue []int

var result int

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

func (q queue) Size() int {
	return len(q)
}

var metro []metroLine

type metroLine struct {
	stations   map[int]bool
	crossLines []int
	color      int
}

func newMetroLine() metroLine {
	ml := metroLine{}
	ml.stations = make(map[int]bool)
	ml.crossLines = make([]int, 0, 0)
	return ml
}

func count(start int, end int) {
	q := queue{}
	for i, v := range metro {
		if v.stations[start] == true {
			q = q.Push(i)
		}
	}
	for q.Size() > 0 {
		var lineNum int
		q, lineNum, _ = q.Pop()

		//if
		fmt.Println(lineNum)
	}

}

func crossLines() {
	for i := 0; i < len(metro)-1; i++ {
		for j := i + 1; j < len(metro); j++ {
			for k, _ := range metro[i].stations {
				if metro[j].stations[k] == true {
					fmt.Printf(" st cross = %d \n", k)
					metro[i].crossLines = append(metro[i].crossLines, j)
					metro[j].crossLines = append(metro[j].crossLines, i)

				}
			}
		}
	}
}

func DoInput(in io.Reader, out io.Writer) {
	var stationCount, metroLineCount int
	fmt.Fscan(in, &stationCount)
	fmt.Fscan(in, &metroLineCount)
	metro = make([]metroLine, metroLineCount, metroLineCount)
	for i := 0; i < metroLineCount; i++ {
		metro[i] = newMetroLine()
		var lineStationCount int
		fmt.Fscan(in, &lineStationCount)
		for j := 0; j < lineStationCount; j++ {
			var station int
			fmt.Fscan(in, &station)
			metro[i].stations[station] = true
		}
	}
	var start, end int
	fmt.Fscan(in, &start, &end)
	crossLines()
	count(start, end)
	fmt.Fprintln(out, metro)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	DoInput(in, out)
}
