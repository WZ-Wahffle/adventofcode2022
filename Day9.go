package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type pos struct {
	x int
	y int
}

func areTouching(a, b pos) bool {
	dist := math.Sqrt(math.Pow(float64(a.x)-float64(b.x), 2) + math.Pow(float64(a.y)-float64(b.y), 2))
	if dist > math.Sqrt(2) {
		return false
	} else {
		return true
	}
}

func moveInDir(a *pos, d rune) {
	switch d {
	case 'U':
		a.y++
		break
	case 'D':
		a.y--
		break
	case 'R':
		a.x++
		break
	case 'L':
		a.x--
		break
	}
}

func catchUp(posH, posT *pos) {
	if !areTouching(*posH, *posT) {
		if posH.x > posT.x {
			moveInDir(posT, 'R')
		}
		if posH.x < posT.x {
			moveInDir(posT, 'L')
		}
		if posH.y > posT.y {
			moveInDir(posT, 'U')
		}
		if posH.y < posT.y {
			moveInDir(posT, 'D')
		}
	}
}

func main() {
	dat, _ := os.ReadFile("input/Day9Input.txt")
	file := strings.Split(string(dat), "\n")

	rope := make([]pos, 10)
	for i := range rope {
		rope[i] = pos{0, 0}
	}

	historyOne, historyTwo := make(map[pos]bool), make(map[pos]bool)
	historyOne[rope[1]] = true
	historyTwo[rope[9]] = true

	for _, i := range file {
		dir := ' '
		length := 0
		_, _ = fmt.Sscanf(i, "%c %d", &dir, &length)

		for range length {
			moveInDir(&rope[0], dir)

			for j := 1; j < 10; j++ {
				catchUp(&rope[j-1], &rope[j])
			}

			historyOne[rope[1]] = true
			historyTwo[rope[9]] = true
		}
	}

	fmt.Println("Part 1:", len(historyOne))
	fmt.Println("Part 2:", len(historyTwo))
}
