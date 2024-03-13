package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
)

type d12pos struct {
	x int
	y int
}

type d12cacheObj struct {
	x     int
	y     int
	steps int
}

// true signifies that branch has been searched
var cache = make(map[d12cacheObj]bool)

func forkPaths(hMap [][]byte, steps int, pos, dest d12pos) {
	if cache[d12cacheObj{pos.x, pos.y, steps}] == true {
		return
	}
	if pos.x == dest.x && pos.y == dest.y {
		fmt.Println("Found!", steps)
		os.Exit(0)
	}
	if steps > 0 {
		if pos.x > 0 && math.Abs(float64(hMap[pos.x-1][pos.y]-hMap[pos.x][pos.y])) <= 1 {
			if !cache[d12cacheObj{pos.x - 1, pos.y, steps - 1}] {
				forkPaths(hMap, steps-1, d12pos{pos.x - 1, pos.y}, dest)
				cache[d12cacheObj{pos.x - 1, pos.y, steps - 1}] = true
			}
		}
		if pos.x < len(hMap)-1 && math.Abs(float64(hMap[pos.x+1][pos.y]-hMap[pos.x][pos.y])) <= 1 {
			if !cache[d12cacheObj{pos.x + 1, pos.y, steps - 1}] {
				forkPaths(hMap, steps-1, d12pos{pos.x + 1, pos.y}, dest)
				cache[d12cacheObj{pos.x + 1, pos.y, steps - 1}] = true
			}
		}
		if pos.y > 0 && math.Abs(float64(hMap[pos.x][pos.y-1]-hMap[pos.x][pos.y])) <= 1 {
			if !cache[d12cacheObj{pos.x, pos.y - 1, steps - 1}] {
				forkPaths(hMap, steps-1, d12pos{pos.x, pos.y - 1}, dest)
				cache[d12cacheObj{pos.x, pos.y - 1, steps - 1}] = true
			}
		}
		if pos.y < len(hMap[0])-1 && math.Abs(float64(hMap[pos.y+1][pos.y]-hMap[pos.x][pos.y])) <= 1 {
			if !cache[d12cacheObj{pos.x, pos.y + 1, steps - 1}] {
				forkPaths(hMap, steps-1, d12pos{pos.x, pos.y + 1}, dest)
				cache[d12cacheObj{pos.x, pos.y + 1, steps - 1}] = true
			}
		}
	}

	cache[d12cacheObj{pos.x, pos.y, steps}] = true

}

func main() {
	dat, _ := os.ReadFile("input/Test.txt")
	file := bytes.Split(dat, []byte("\n"))

	start, end := d12pos{0, 0}, d12pos{0, 0}

	for ii, i := range file {
		for jj, j := range i {
			if j == 'S' {
				start = d12pos{ii, jj}
				file[ii][jj] = 'a'
			}
			if j == 'E' {
				end = d12pos{ii, jj}
				file[ii][jj] = 'z'
			}
		}
	}

	maxDist := len(file) * len(file[0])
	for i := range maxDist {
		fmt.Println(i)
		forkPaths(file, i, start, end)
		if i == 31 {
			for j := range cache {
				fmt.Println(j.x, j.y)

			}
		}
		cache = make(map[d12cacheObj]bool)
	}

}
