package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type fsType int32

const (
	FILE   fsType = 0
	FOLDER fsType = 1
)

type fsObj struct {
	kind     fsType
	content  []fsObj
	filesize int32
	name     string
	parent   *fsObj
}

func (f fsObj) size() int32 {
	if f.kind == FILE {
		return f.filesize
	} else {
		sum := int32(0)
		for _, i := range f.content {
			sum += i.size()
		}
		return sum
	}
}

func (f fsObj) getLessThan(dirs *[]int32, edge int32) {
	if f.kind == FOLDER {
		if f.size() <= edge {
			*dirs = append(*dirs, f.size())
		}
		for _, i := range f.content {
			i.getLessThan(dirs, edge)
		}
	}
}

func main() {
	dat, _ := os.ReadFile("input/Day7Input.txt")
	data := strings.Split(string(dat), "\n")

	root := fsObj{
		kind:     FOLDER,
		content:  make([]fsObj, 0),
		filesize: 0,
		name:     "/",
		parent:   nil,
	}

	current := &root

	filling := false

	for _, i := range data {
		if i == "$ cd /" { // ignore first line, always the same
			continue
		}

		if i[0] == '$' {
			// command parsing
			if i[2:4] == "ls" {
				filling = true
				continue
			} else {
				filling = false
				target := i[5:]

				if target == ".." {
					current = current.parent
					continue
				}

				for j := range current.content {
					if current.content[j].name == target {

						current = &current.content[j]
						break
					}
				}
			}
		}

		if filling {
			if i[:3] == "dir" {
				current.content = append(current.content, fsObj{
					kind:     FOLDER,
					content:  make([]fsObj, 0),
					filesize: 0,
					name:     i[4:],
					parent:   current,
				})
			} else {
				sz, nm := int32(0), ""
				_, _ = fmt.Sscanf(i, "%d %s", &sz, &nm)
				current.content = append(current.content, fsObj{
					kind:     FILE,
					content:  nil,
					filesize: sz,
					name:     nm,
					parent:   current,
				})
			}
		}
	}

	dirsLess100k := make([]int32, 0)
	dirs := make([]int32, 0)

	root.getLessThan(&dirsLess100k, 100000)
	root.getLessThan(&dirs, math.MaxInt32)

	total := int32(0)
	lowestPossible := int32(math.MaxInt32)

	for _, i := range dirsLess100k {
		total += i
	}
	for _, i := range dirs {
		if i < lowestPossible && i >= 30000000-(70000000-root.size()) {
			lowestPossible = i
		}
	}

	fmt.Println("Part 1:", total)
	fmt.Println("Part 2:", lowestPossible)
}
