package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	dat, err := os.ReadFile("./input/Day1Input.txt")
	if err != nil {
		panic(err)
	}
	file := bytes.Split(dat, []byte("\n"))

	temp := 0
	var out []int

	for i := range file {
		if len(file[i]) == 0 {
			out = append(out, temp)
			temp = 0
		} else {
			parseInt, _ := strconv.Atoi(string(file[i]))
			temp += parseInt
		}
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i] > out[j]
	})

	fmt.Println("Part 1: ", out[0])
	fmt.Println("Part 2: ", out[0]+out[1]+out[2])
}
