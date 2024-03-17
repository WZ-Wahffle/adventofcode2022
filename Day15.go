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
	positives := make([]int, 0)
	negatives := make([]int, 0)

	p1Line := 10
	minX, maxX := math.MaxInt32, 0

	for _, i := range file {
		sx, dx, sy, dy := 0, 0, 0, 0
		_, _ = fmt.Sscanf(i, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &dx, &dy)
		distance := int(math.Abs(float64(sx-dx)) + math.Abs(float64(sy-dy)))

		minX = min(minX, sx-dx)
		maxX = max(maxX, sx+dx)
		ranges[d15Point{sx, sy}] = distance
		beacons = append(beacons, d15Point{dx, dy})
		positives = append(positives, sy-sx-distance-1)
		negatives = append(negatives, sy+sx-distance-1)
	}

	cannot := 0
	for i := minX; i < maxX; i++ {
		for j, jj := range ranges {
			if int(math.Abs(float64(i-j.x))+math.Abs(float64(p1Line-j.y))) <= jj && !slices.Contains(beacons, d15Point{i, p1Line}) {
				cannot++
				break
			}
		}
	}
	fmt.Println("Part 1:", cannot)

	for _, i := range positives {
		for _, j := range negatives {
			valid := true
			interX := (j - i) / 2
			interY := (j + i) / 2
			for k, kk := range ranges {
				if int(math.Abs(float64(interX-k.x))+math.Abs(float64(interY-k.y))) <= kk || slices.Contains(beacons, d15Point{interX, interY}) {
					valid = false
					break
				}
			}

			if valid && interX >= 0 && interX <= 4000000 && interY >= 0 && interY <= 4000000 {
				fmt.Println("Part 2:", int64(interX)*int64(4000000)+int64(interY))
				os.Exit(0)
			}
		}
	}
}
