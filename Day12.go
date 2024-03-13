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

type dir int

const (
	NONE  = 0
	UP    = 1
	DOWN  = 2
	LEFT  = 3
	RIGHT = 4
)

// true signifies that branch has been searched
var cache = make(map[d12cacheObj]bool)
var finished = -1

func forkPaths(hMap [][]byte, steps, initSteps int, pos, dest d12pos, prevDir dir) {
	if cache[d12cacheObj{pos.x, pos.y, steps}] == true || finished != -1 {
		return
	}

	if pos.x == dest.x && pos.y == dest.y {
		finished = initSteps
		return
	}
	if steps > 0 {
		// cur = s next = q
		if prevDir != RIGHT && pos.x > 0 && int(hMap[pos.x-1][pos.y])-int(hMap[pos.x][pos.y]) <= 1 {
			if !cache[d12cacheObj{pos.x - 1, pos.y, steps - 1}] {
				forkPaths(hMap, steps-1, initSteps, d12pos{pos.x - 1, pos.y}, dest, LEFT)
				cache[d12cacheObj{pos.x - 1, pos.y, steps - 1}] = true
			}
		}
		if prevDir != LEFT && pos.x < len(hMap)-1 && int(hMap[pos.x+1][pos.y])-int(hMap[pos.x][pos.y]) <= 1 {
			if !cache[d12cacheObj{pos.x + 1, pos.y, steps - 1}] {
				forkPaths(hMap, steps-1, initSteps, d12pos{pos.x + 1, pos.y}, dest, RIGHT)
				cache[d12cacheObj{pos.x + 1, pos.y, steps - 1}] = true
			}
		}
		if prevDir != DOWN && pos.y > 0 && int(hMap[pos.x][pos.y-1])-int(hMap[pos.x][pos.y]) <= 1 {
			if !cache[d12cacheObj{pos.x, pos.y - 1, steps - 1}] {
				forkPaths(hMap, steps-1, initSteps, d12pos{pos.x, pos.y - 1}, dest, UP)
				cache[d12cacheObj{pos.x, pos.y - 1, steps - 1}] = true
			}
		}
		if prevDir != UP && pos.y < len(hMap[0])-1 && int(hMap[pos.x][pos.y+1])-int(hMap[pos.x][pos.y]) <= 1 {
			if !cache[d12cacheObj{pos.x, pos.y + 1, steps - 1}] {
				forkPaths(hMap, steps-1, initSteps, d12pos{pos.x, pos.y + 1}, dest, DOWN)
				cache[d12cacheObj{pos.x, pos.y + 1, steps - 1}] = true
			}
		}
	}

	cache[d12cacheObj{pos.x, pos.y, steps}] = true

}

func main() {
	dat, _ := os.ReadFile("input/Day12Input.txt")
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
		forkPaths(file, i, i, start, end, NONE)
		if finished != -1 {
			break
		}
	}

	fmt.Println("Part 1:", finished)
	finished = -1
	minFinished := math.MaxInt

	for i, ii := range file {
		for j, jj := range ii {
			if jj == 'a' {
				cache = make(map[d12cacheObj]bool)
				for k := range maxDist {
					forkPaths(file, k, k, d12pos{i, j}, end, NONE)
					if finished != -1 {
						break
					}
				}
				if finished != -1 {
					minFinished = min(minFinished, finished)
					finished = -1
				}
			}
		}
	}

	fmt.Println("Part 2:", minFinished)
}
