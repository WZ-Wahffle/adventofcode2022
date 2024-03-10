package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("./input/Day5Input.txt")
	file := strings.Split(string(dat), "\n")

	passedPile := false
	crates := make([]string, 0)

	for _, i := range file {
		if !passedPile {

			for len(crates) < (len(i)+1)/4 {
				crates = append(crates, "")
			}

			if len(i) == 0 {
				passedPile = true
				continue
			}
			for j := 0; j < len(i); j += 4 {
				if i[j] == '[' {

					crates[j/4] = string(append([]byte(crates[j/4]), i[j+1]))
				}
			}
		} else {
			cnt, from, to := 0, 0, 0
			_, err := fmt.Sscanf(i, "move %d from %d to %d", &cnt, &from, &to)
			if err != nil {
				panic(err)
			}
			from--
			to--

			for ; cnt > 0; cnt-- {
				crates[to] = string(append([]byte{crates[from][0]}, []byte(crates[to])...))
				crates[from] = crates[from][1:]
			}
		}
	}

	out := ""

	for _, i := range crates {
		out = string(append([]byte(out), i[0]))
	}

	fmt.Println("Part 1:", out)

	passedPile = false
	crates = make([]string, 0)

	for _, i := range file {
		if !passedPile {

			for len(crates) < (len(i)+1)/4 {
				crates = append(crates, "")
			}

			if len(i) == 0 {
				passedPile = true
				continue
			}
			for j := 0; j < len(i); j += 4 {
				if i[j] == '[' {

					crates[j/4] = string(append([]byte(crates[j/4]), i[j+1]))
				}
			}
		} else {
			cnt, from, to := 0, 0, 0
			_, err := fmt.Sscanf(i, "move %d from %d to %d", &cnt, &from, &to)
			if err != nil {
				panic(err)
			}
			from--
			to--

			crates[to] = string(append([]byte(crates[from][:cnt]), []byte(crates[to])...))
			crates[from] = crates[from][cnt:]

		}
	}

	out = ""

	for _, i := range crates {
		out = string(append([]byte(out), i[0]))
	}

	fmt.Println("Part 1:", out)
}
