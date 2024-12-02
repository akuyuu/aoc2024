package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partOne() {
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(f)
	sum := 0
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			break
		}
		ls := strings.Split(string(line), " ")
		if processOne(ls) {
			sum += 1
		}
	}

	fmt.Println(sum)
}

func partTwo() {
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(f)
	sum := 0
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			break
		}
		ls := strings.Split(string(line), " ")
		if processTwo(ls) <= 1 {
			sum += 1
		}
	}
	fmt.Println(sum)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func sig(n int) int {
	if n == 0 {
		return 0
	}
	return n / abs(n)
}

func processOne(ls []string) bool {
	ln := make([]int, len(ls), len(ls))
	for i, j := range ls {
		ln[i], _ = strconv.Atoi(j)
	}
	sigd := sig(ln[1] - ln[0])
	for i := 1; i < len(ln); i++ {
		dif := ln[i] - ln[i-1]
		if abs(dif) == 0 || abs(dif) > 3 {
			return false
		}
		if sig(dif) != sigd {
			return false
		}

	}
	return true
}

func processTwo(ls []string) int {
	ln := make([]int, len(ls), len(ls))
	out := 0
	for i, j := range ls {
		ln[i], _ = strconv.Atoi(j)
	}
	sigd := sig(ln[1] - ln[0])
	for i := 1; i < len(ln); i++ {
		dif := ln[i] - ln[i-1]
		if abs(dif) == 0 || abs(dif) > 3 {
			out += 1
		} else if sig(dif) != sigd {
			out += 1
		}

	}
	return out
}

func main() {
	partOne()
	partTwo()
}
