package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func parseInts(ipAddr string, buff *[4][3]byte) (int, int, int, int) {
	num := 0
	i := 0
	for _, v := range ipAddr {
		if v == '.' {
			if i == 1 {
				buff[num][2] = buff[num][0]
				buff[num][1] = '0'
				buff[num][0] = '0'
			} else if i == 2 {
				buff[num][2] = buff[num][1]
				buff[num][1] = buff[num][0]
				buff[num][0] = '0'
			}

			num++
			i = 0
			continue
		}

		buff[num][i] = byte(v)
		i++
	}

	if i == 1 {
		buff[num][2] = buff[num][0]
		buff[num][1] = '0'
		buff[num][0] = '0'
	} else if i == 2 {
		buff[num][2] = buff[num][1]
		buff[num][1] = buff[num][0]
		buff[num][0] = '0'
	}

	a, err := strconv.Atoi(string(buff[0][:]))
	if err != nil {
		panic(err)
	}

	b, err := strconv.Atoi(string(buff[1][:]))
	if err != nil {
		panic(err)
	}

	c, err := strconv.Atoi(string(buff[2][:]))
	if err != nil {
		panic(err)
	}

	d, err := strconv.Atoi(string(buff[3][:]))
	if err != nil {
		panic(err)
	}

	return a, b, c, d
}
