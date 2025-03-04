package main

import (
	"bufio"
	"fmt"
	"os"
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

	arr := make([][][][256]bool, 256)

	file, err := os.Open("ip_addresses")
	if err != nil {
		panic(fmt.Sprint("ERROR reading file:", err))
	}

	defer file.Close()

	ParseIPBuff := [4][3]byte{}

	// TODO try custom reader - to have less allocations
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		a, b, c, d := parseInts(line, &ParseIPBuff)
		if arr[a] == nil {
			arr[a] = make([][][256]bool, 256)
		}
		if arr[a][b] == nil {
			arr[a][b] = make([][256]bool, 256)
		}
		if arr[a][b][c][d] {
			continue
		}
		count++
		arr[a][b][c][d] = true
	}

	fmt.Println("THE ANSWER: ", count)
}

// We have this function against strings.Split(line, ".")
// to have less allocations.
//
// ipAddr: 123.456.78.9
// buff:
// [[0, 0, 0], [0, 0, 0], [0, 0, 0], [0, 0, 0]]
// [[1, 0, 0], [0, 0, 0], [0, 0, 0], [0, 0, 0]]
// [[1, 2, 0], [0, 0, 0], [0, 0, 0], [0, 0, 0]]
// [[1, 2, 3], [0, 0, 0], [0, 0, 0], [0, 0, 0]]
// [[1, 2, 3], [4, 0, 0], [0, 0, 0], [0, 0, 0]]
// [[1, 2, 3], [4, 5, 0], [0, 0, 0], [0, 0, 0]]
// [[1, 2, 3], [4, 5, 6], [0, 0, 0], [0, 0, 0]]
// [[1, 2, 3], [4, 5, 6], [7, 0, 0], [0, 0, 0]]
// [[1, 2, 3], [4, 5, 6], [7, 8, 0], [0, 0, 0]]
// [[1, 2, 3], [4, 5, 6], [0, 7, 8], [0, 0, 0]]
// [[1, 2, 3], [4, 5, 6], [0, 7, 8], [9, 0, 0]]
// [[1, 2, 3], [4, 5, 6], [0, 7, 8], [0, 0, 9]]
func parseInts(ipAddr string, buff *[4][3]byte) (byte, byte, byte, byte) {
	for i, j, I := 0, 0, 0; I <= len(ipAddr); I++ {
		if I == len(ipAddr) || ipAddr[I] == '.' {
			// Move [1, 0, 0] to [0, 0, 1]
			if j == 1 {
				buff[i][2] = buff[i][0]
				buff[i][1] = '0'
				buff[i][0] = '0'
			} else if j == 2 {
				buff[i][2] = buff[i][1]
				buff[i][1] = buff[i][0]
				buff[i][0] = '0'
			}

			i++
			j = 0
			continue
		}

		buff[i][j] = byte(ipAddr[I])
		j++
	}

	return byteMerge(buff[0]), byteMerge(buff[1]), byteMerge(buff[2]), byteMerge(buff[3])
}

func byteMerge(b [3]byte) (res byte) {
	res += b[2]
	res += b[1] * 10
	res += b[0] * 100

	return
}
