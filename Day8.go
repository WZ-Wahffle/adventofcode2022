package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	dat, _ := os.ReadFile("./input/Day8Input.txt")
	file := bytes.Split(dat, []byte("\n"))

	visible := make(map[int]bool)

	// left, right
	for i := range file {
		largestLeft, largestRight := 0, 0
		for j := range file[i] {
			leftVisIdx, rightVisIdx := i*len(file[i])+j, i*len(file[i])+len(file[i])-j-1
			if largestLeft < int(file[i][j]) {
				largestLeft = int(file[i][j])
				visible[leftVisIdx] = true
			}
			if largestRight < int(file[i][len(file[i])-j-1]) {
				largestRight = int(file[i][len(file[i])-j-1])
				visible[rightVisIdx] = true
			}
		}
	}

	// up, down
	for j := range file[0] {
		largestUp, largestDown := 0, 0
		for i := range file {
			upVisIdx, downVisIdx := i*len(file[i])+j, (len(file)-i-1)*len(file[i])+j
			if largestUp < int(file[i][j]) {
				largestUp = int(file[i][j])
				visible[upVisIdx] = true
			}
			if largestDown < int(file[len(file)-i-1][j]) {
				largestDown = int(file[len(file)-i-1][j])
				visible[downVisIdx] = true
			}
		}
	}

	highestScenic := 0

	for i := range file {
		for j := range file[i] {
			l, r, u, d := 0, 0, 0, 0
			// down
			for x := 1; x+i < len(file); x++ {
				d++
				if file[i+x][j] >= file[i][j] {
					break
				}
			}
			// up
			for x := -1; x+i >= 0; x-- {
				u++
				if file[i+x][j] >= file[i][j] {
					break
				}
			}
			// right
			for x := 1; x+j < len(file[i]); x++ {
				r++
				if file[i][j+x] >= file[i][j] {
					break
				}
			}
			// left
			for x := -1; x+j >= 0; x-- {
				l++
				if file[i][j+x] >= file[i][j] {
					break
				}
			}

			highestScenic = max(highestScenic, u*d*l*r)
		}
	}

	fmt.Println("Part 1:", len(visible))
	fmt.Println("Part 2:", highestScenic)
}
