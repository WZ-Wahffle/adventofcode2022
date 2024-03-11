package main

import (
	"fmt"
	"os"
)

func main() {
	dat, _ := os.ReadFile("./input/Day6Input.txt")

	for i := 0; i < len(dat)-4; i++ {
		counts := make(map[byte]bool)
		for j := range 4 {
			counts[dat[j+i]] = true
		}
		if len(counts) == 4 {
			fmt.Println("Part 1:", i+4)
			break
		}
	}

	for i := 0; i < len(dat)-14; i++ {
		counts := make(map[byte]bool)
		for j := range 14 {
			counts[dat[j+i]] = true
		}
		if len(counts) == 14 {
			fmt.Println("Part 2:", i+14)
			break
		}
	}
}
