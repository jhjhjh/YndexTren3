package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type colector struct {
	number int
	count  int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var diegoStickerCount int
	fmt.Fscan(in, &diegoStickerCount)
	stickers := make(map[int]bool)
	var sticker int
	for i := 0; i < diegoStickerCount; i++ {
		fmt.Fscan(in, &sticker)
		stickers[sticker] = true
	}

	var colCount int
	fmt.Fscan(in, &colCount)
	colectors := make([]colector, colCount)
	var col int
	for i := 0; i < colCount; i++ {
		fmt.Fscan(in, &col)
		colectors[i].number = i
		colectors[i].count = col
	}
	if diegoStickerCount == 0 {
		for i := 0; i < len(colectors); i++ {
			fmt.Fprintln(out, 0)
		}

	}
	stickerUniq := make([]int, len(stickers))
	i := 0
	for st, _ := range stickers {
		stickerUniq[i] = st
		i++
	}
	sort.Slice(stickerUniq, func(i, j int) bool { return stickerUniq[i] < stickerUniq[j] })
	sort.Slice(colectors, func(i, j int) bool { return colectors[i].count < colectors[j].count })
	result := make([]int, len(colectors))
	j := 0
	for i := 0; i < len(stickerUniq); {
		if i+1 == len(stickerUniq) {
			if stickerUniq[i] < colectors[j].count {
				result[colectors[j].number] = i + 1
			}
			j++
			if j == len(colectors) {
				break
			}
		} else {
			if stickerUniq[i] > colectors[j].count {
				j++
			} else {
				if colectors[j].count < stickerUniq[i+1] {
					result[colectors[j].number] = i + 1
					j++
				} else {
					i++
				}
			}
			if j == len(colectors) {
				break
			}
		}
	}
	for _, v := range result {
		fmt.Fprintln(out, v)
	}

}
