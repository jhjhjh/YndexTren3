package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"io"
	"os"
	"runtime"
)

func printRing(r *ring.Ring) {
	t := r
	for i := 0; i < 4; i++ {
		fmt.Println(t.Value)
		t = t.Next()
	}
}

func printPtr(fst *ring.Ring, scd *ring.Ring, thrd *ring.Ring, fth *ring.Ring) {
	fmt.Println(fst.Value)
	fmt.Println(scd.Value)
	fmt.Println(thrd.Value)
	fmt.Println(fth.Value)

}

func minVar(in []uint) uint {
	min := in[0]
	for _, v := range in {
		if v < min {
			min = v
		}
	}
	return min
}

func Compute(in <-chan []uint, out chan<- uint) {
	initStruct := []uint{4000, 4000, 4000, 0}
	myRing := ring.New(4)
	for i := 0; i < 4; i++ {
		myRing.Value = initStruct
		myRing = myRing.Next()
	}
	initStruct = nil
	fst := myRing
	myRing = myRing.Next()
	scd := myRing
	myRing = myRing.Next()
	thrd := myRing
	myRing = myRing.Next()
	fth := myRing
	myRing = myRing.Next()
	for tmp := range in {
		fst.Value = tmp
		fth = fth.Next()
		thrd = thrd.Next()
		scd = scd.Next()
		fst = fst.Next()
		vars := make([]uint, 3, 3)

		vars[0] = fth.Value.([]uint)[0] + thrd.Value.([]uint)[3]
		vars[1] = thrd.Value.([]uint)[1] + scd.Value.([]uint)[3]
		vars[2] = scd.Value.([]uint)[2] + fst.Value.([]uint)[3]
		fth.Value.([]uint)[3] = minVar(vars)
		//printRing(fst)
		//fmt.Println()

		//printPtr(fst, scd, thrd, fth)
		//fmt.Println()
	}
	out <- fth.Value.([]uint)[3]
}

func Do(in io.Reader, out io.Writer) {
	var n uint
	fmt.Fscan(in, &n)
	dataChan := make(chan []uint)
	result := make(chan uint)
	go Compute(dataChan, result)
	for i := uint(0); i < n; i++ {
		dataLine := make([]uint, 4, 4)
		fmt.Fscan(in, &dataLine[0], &dataLine[1], &dataLine[2])
		dataChan <- dataLine
	}
	close(dataChan)
	fmt.Fprint(out, <-result)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Do(in, out)
}
