package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("./input/Day4Input.txt")
	file := strings.Split(string(dat), "\n")

	p1Count, p2Count := 0, 0

	for _, i := range file {
		first, second, third, fourth := 0, 0, 0, 0
		_, err := fmt.Sscanf(i, "%d-%d,%d-%d", &first, &second, &third, &fourth)
		if err != nil {
			panic(err)
		}

		if (first >= third && second <= fourth) || (first <= third && second >= fourth) {
			p1Count++
		}

		if (second-first)+(fourth-third) >= max(first, second, third, fourth)-min(first, second, third, fourth) {
			p2Count++
		}
	}

	fmt.Println("Part 1:", p1Count)
	fmt.Println("Part 2:", p2Count)
}
