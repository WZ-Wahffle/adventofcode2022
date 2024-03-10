package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	dat, _ := os.ReadFile("./input/Day2Input.txt")
	file := bytes.Split(dat, []byte("\n"))
	score := 0

	for i := range file {
		if len(file[i]) == 0 {
			continue
		}
		first := file[i][0]
		second := file[i][2]

		if first == second-23 {
			score += 3
		} else if (first == 'A' && second == 'Y') || (first == 'B' && second == 'Z') || (first == 'C' && second == 'X') {
			score += 6
		}
		score += int(second - 23 - 64)
	}

	fmt.Println("Part 1: ", score)

	score = 0

	for i := range file {
		if len(file[i]) == 0 {
			continue
		}
		first := file[i][0]
		second := file[i][2]

		switch second {
		case 'X':
			if first != 'A' {
				score += int(first-64) - 1
			} else {
				score += 3
			}
			break
		case 'Y':
			score += 3 + int(first-64)
			break
		case 'Z':
			score += 6
			if first != 'C' {
				score += int(first-64) + 1
			} else {
				score += 1
			}
			break
		}
	}

	fmt.Println("Part 2: ", score)
}
