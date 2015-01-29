package main

import (
	"fmt"
	"time"

	"bufio"
	"log"
	"os"

	//"sort"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func merge_sort(l []int) []int {
	if len(l) <= 1 {
		return l
	}
	mid := len(l) / 2
	left := merge_sort(l[:mid])
	right := merge_sort(l[mid:])
	//fmt.Println(left, right)
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
	//fmt.Println(result)
	return result
}

func main() {
	l := make([]int, 1000000)

	lines, err := readLines("arr.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for i, line := range lines {
		l[i], _ = strconv.Atoi(line)
	}

	start := time.Now()
	//sort.Ints(l)
	l = merge_sort(l)
	end := time.Since(start)
	fmt.Println(end)
	//fmt.Println(l)
}
