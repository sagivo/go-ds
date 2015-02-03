package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
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

func merge_sort(l []int, c chan []int) {
	if len(l) < 2 {
		c <- l
	}
	mid := len(l) / 2
	go merge_sort(l[:mid], c)
	go merge_sort(l[mid:], c)
	c <- merge(<-c, <-c)
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
	//sort.Ints(l)
	c := make(chan []int)
	merge_sort(l, c)
	l = <-c
	end := time.Since(start)
	fmt.Println(end)
}
