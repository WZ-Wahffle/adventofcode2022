package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type d15Point struct {
	x int
	y int
}

func main() {
	dat, _ := os.ReadFile("input/Day15Input.txt")
	file := strings.Split(string(dat), "\n")
	ranges := make(map[d15Point]int)
	beacons := make([]d15Point, 0)

	p1Line := 2000000
	minX, maxX := math.MaxInt32, 0

	for _, i := range file {
		sx, dx, sy, dy := 0, 0, 0, 0
		_, _ = fmt.Sscanf(i, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &dx, &dy)
		distance := int(math.Abs(float64(sx-dx)) + math.Abs(float64(sy-dy)))

		minX = min(minX, sx-dx)
		maxX = max(maxX, sx+dx)
		ranges[d15Point{sx, sy}] = distance
		beacons = append(beacons, d15Point{dx, dy})
	}

	cannot := 0
	for i := minX; i < maxX; i++ {
		found := false
		for j, jj := range ranges {
			if int(math.Abs(float64(i-j.x))+math.Abs(float64(p1Line-j.y))) <= jj && !slices.Contains(beacons, d15Point{i, p1Line}) {
				cannot++
				found = true
				break
			}
		}
		if !found && i >= 0 && i <= 4000000 {
			fmt.Println("Part 2:", i, p1Line, int64(i)*int64(4000000)+int64(p1Line))
		}
	}
	fmt.Println("Part 1:", cannot)
}
