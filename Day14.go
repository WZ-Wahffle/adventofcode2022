package main

import (
	"fmt"
	"os"
	"strings"
)

type d14Pair struct {
	x int
	y int
}

func d14DropSand(maxY int, sandMap [][]uint8) bool {
	pos := d14Pair{500, 0}
	for pos.y < maxY {
		if sandMap[pos.y+1][pos.x] != '#' {
			pos.y++
		} else if sandMap[pos.y+1][pos.x-1] != '#' {
			pos.x--
			pos.y++
		} else if sandMap[pos.y+1][pos.x+1] != '#' {
			pos.x++
			pos.y++
		} else {
			sandMap[pos.y][pos.x] = '#'
			return true
		}
	}

	return false
}

func main() {
	dat, _ := os.ReadFile("input/Day14Input.txt")
	file := strings.Split(string(dat), "\n")

	sandMap := make([][]uint8, 300)
	for j := range sandMap {
		sandMap[j] = make([]uint8, 1000)
	}

	maxY := 0
	for _, i := range file {
		pairs := strings.Split(i, " -> ")
		points := make([]d14Pair, 0)
		for _, j := range pairs {
			points = append(points, d14Pair{})
			_, _ = fmt.Sscanf(j, "%d,%d", &points[len(points)-1].x, &points[len(points)-1].y)
			maxY = max(points[len(points)-1].y+2, maxY)
		}

		for j := 1; j < len(points); j++ {
			first, second := points[j-1], points[j]

			for first != second {
				sandMap[first.y][first.x] = '#'
				if first.x < second.x {
					first.x++
				}
				if first.x > second.x {
					first.x--
				}
				if first.y < second.y {
					first.y++
				}
				if first.y > second.y {
					first.y--
				}
			}

			sandMap[second.y][second.x] = '#'
		}
	}

	cnt := 0
	for {
		if !d14DropSand(maxY, sandMap) {
			break
		}
		cnt++
	}

	fmt.Println("Part 1:", cnt)

	for i := range sandMap[0] {
		sandMap[maxY][i] = '#'
	}

	for {
		cnt++
		d14DropSand(maxY, sandMap)
		if sandMap[0][500] == '#' {
			break
		}
	}
	fmt.Println("Part 2:", cnt)
}
