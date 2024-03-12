package main

import (
	"fmt"
	"os"
	"strings"
)

var cycles = 0
var x = 1
var signalStrength = 0
var part2Output = "\n"

func incCycle() {

	if cycles%40 >= x%40-1 && cycles%40 <= x%40+1 {
		part2Output = string(append([]byte(part2Output), '#'))
	} else {
		part2Output = string(append([]byte(part2Output), '.'))
	}

	cycles++
	if cycles%40 == 0 {
		part2Output = string(append([]byte(part2Output), '\n'))
	}
	if cycles%40 == 20 && cycles <= 220 {
		signalStrength += x * cycles
	}
}

func main() {
	dat, _ := os.ReadFile("input/Day10Input.txt")
	file := strings.Split(string(dat), "\n")

	for _, i := range file {

		if i[0] == 'n' {
			incCycle()
		} else {
			val := 0
			_, _ = fmt.Sscanf(i, "addx %d", &val)
			incCycle()
			incCycle()
			x += val
		}
	}

	fmt.Println("Part 1:", signalStrength)
	fmt.Println("Part 2:", part2Output)
}
