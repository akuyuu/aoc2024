package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	li := make([]int, 1000, 1000)
	ri := make([]int, 1000, 1000)
	sum := 0
	b := make([]byte, 14)
	// 5 digits + 3 spaces + 5 digits + \n
	for i := 0; i < 1000; i++ {
		_, err := f.Read(b)
		if err != nil {
			break
		}
		li[i], _ = strconv.Atoi(string(b[:5]))
		ri[i], _ = strconv.Atoi(string(b[8:13]))
	}
	slices.Sort(li)
	slices.Sort(ri)
	for i := 0; i < 1000; i++ {
		sum += abs(li[i] - ri[i])
	}
	fmt.Println(sum)
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}
