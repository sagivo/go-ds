package main

import (
	"./fileio.go"
	"fmt"
	"sort"
	"strconv"
	"time"
)

func merge_sort(l []int) []int {
	if len(l) <= 1 {
		return l
	}
	mid := len(l) / 2
	left := merge_sort(l[:mid])
	right := merge_sort(l[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	var i, j, ind int
	result := make([]int, len(left)+len(right))

	for ind = 0; i < len(left) && j < len(right); ind++ {
		if left[i] <= right[j] {
			result[ind] = left[i]
			i++
		} else {
			result[ind] = right[j]
			j++
		}
	}
	for ; i < len(left); i, ind = i+1, ind+1 {
		result[ind] = left[i]
	}
	for ; j < len(right); j, ind = j+1, ind+1 {
		result[ind] = right[j]
	}
	return result
}

func main() {
	l := make([]int, 10000000)

	lines, err := readLines("arr.txt")
	for i, line := range lines {
		if i <= len(l) {
			l[i], err = strconv.Atoi(line)
		}
		if err != nil {
			fmt.Println(err)
		}
	}

	start := time.Now()
	sort.Ints(l)
	//l = merge_sort(l)
	end := time.Since(start)
	fmt.Println(end)
}
