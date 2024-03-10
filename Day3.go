package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("./input/Day3Input.txt")
	file := strings.Split(string(dat), "\n")

	sum := 0

	for _, i := range file {
		duplicates := map[uint8]bool{}
		first := i[:len(i)/2]
		second := i[len(i)/2:]

		for j := range first {
			for k := range second {
				if first[j] == second[k] {
					duplicates[first[j]] = true
				}
			}
		}

		for j := range duplicates {
			if j >= 'a' {
				sum += int(j - 96)
			} else {
				sum += int(j - 38)
			}
		}

	}

	fmt.Println("Part 1: ", sum)

	sum = 0

	for i := 0; i < len(file); i += 3 {
		first := file[i]
		second := file[i+1]
		third := file[i+2]
		duplicates := map[uint8]bool{}

		for j := range first {
			for k := range second {
				for l := range third {
					if first[j] == second[k] && second[k] == third[l] {
						duplicates[first[j]] = true
					}
				}
			}
		}

		for j := range duplicates {
			if j >= 'a' {
				sum += int(j - 96)
			} else {
				sum += int(j - 38)
			}
		}
	}

	fmt.Println("Part 2: ", sum)
}
