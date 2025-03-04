package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// type rang [2]int

// func (r rang) isIncludes(in int) bool {
// 	if r[0] >= in && r[1] <= in {
// 		return true
// 	}
// 	return false
// }

// Ranges example:
// [
//
//	[0-1000]
//	[1005-2000]
//	[2001-3000]
//	...
//
// ]
// type ipv4Ranged []rang

// func (s ipv4Ranged) add(line int) bool {
// 	for _, r := range s {
// 		if r[1] < line
// 		if r.isIncludes(line) {
// 			return true
// 		}
// 	}
// 	// TODO add to s
// }

func main() {
	count := 0

	arr := make([][][][]bool, 256)

	file, err := os.Open("ip_addresses")
	if err != nil {
		panic(fmt.Sprint("ERROR reading file:", err))
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		groups := strings.Split(line, ".")
		a, b, c, d := strToInt(groups[0]), strToInt(groups[1]), strToInt(groups[2]), strToInt(groups[3])
		if arr[a] == nil {
			arr[a] = make([][][]bool, 256)
		}
		if arr[a][b] == nil {
			arr[a][b] = make([][]bool, 256)
		}
		if arr[a][b][c] == nil {
			arr[a][b][c] = make([]bool, 256)
		}
		if arr[a][b][c][d] {
			continue
		}
		count++
		arr[a][b][c][d] = true
	}

	fmt.Println("THE ANSWER: ", count)
}

func strToInt(s string) int {
	in, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return in
}
