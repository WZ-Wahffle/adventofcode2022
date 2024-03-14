package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type d13Composite struct {
	value    int
	contents []d13Composite
	parent   *d13Composite
}

var d13NoSol = false

func fillComposite(content *string, comp *d13Composite) {

	for i := 0; i < len(*content); i++ {
		if (*content)[i] == ',' {
			*content = (*content)[:i] + "x" + (*content)[i+1:]

		} else if (*content)[i] == '[' {
			*content = (*content)[:i] + "x" + (*content)[i+1:]
			comp.contents = append(comp.contents, d13Composite{-1, make([]d13Composite, 0), comp})
			fillComposite(content, &comp.contents[len(comp.contents)-1])
			for i < len(*content) && (*content)[i] == 'x' {
				i++
			}
			i--
		} else if (*content)[i] == ']' {
			*content = (*content)[:i] + "x" + (*content)[i+1:]
			return
		} else if (*content)[i] != 'x' {
			temp, _ := strconv.Atoi(string((*content)[i]))
			for i < len(*content)-1 && unicode.IsDigit(rune((*content)[i+1])) {
				*content = (*content)[:i] + "x" + (*content)[i+1:]
				i++
				temp *= 10
				temp2, _ := strconv.Atoi(string((*content)[i]))
				temp += temp2
			}
			comp.contents = append(comp.contents, d13Composite{temp, nil, comp})
			*content = (*content)[:i] + "x" + (*content)[i+1:]
		}

	}
}

func cmpComposite(c1, c2 d13Composite) bool {

	for i := range max(len(c1.contents), len(c2.contents)) {
		// if two lists and one is shorter
		if i == len(c1.contents) {
			return true
		} else if i == len(c2.contents) {
			return false
		}

		// left is list but right isn't
		if c1.contents[i].value == -1 && c2.contents[i].value != -1 {
			c2.contents[i].contents = append(c2.contents[i].contents, d13Composite{c2.contents[i].value, nil, &c2})
			c2.contents[i].value = -1
		}

		// right is list but left isn't
		if c1.contents[i].value != -1 && c2.contents[i].value == -1 {
			c1.contents[i].contents = append(c1.contents[i].contents, d13Composite{c1.contents[i].value, nil, &c1})
			c1.contents[i].value = -1
		}

		// two lists
		if c1.contents[i].value == -1 && c2.contents[i].value == -1 {

			ret := cmpComposite(c1.contents[i], c2.contents[i])
			if d13NoSol {
				d13NoSol = false
				continue
			} else {
				return ret
			}
		}

		// both integers
		if c1.contents[i].value != -1 && c2.contents[i].value != -1 {
			if c1.contents[i].value > c2.contents[i].value {
				return false
			} else if c1.contents[i].value < c2.contents[i].value {
				return true
			} else {
				continue
			}
		}

	}

	d13NoSol = true
	return true
}

func main() {
	dat, _ := os.ReadFile("input/Day13Input.txt")
	file := strings.Split(string(dat), "\n")
	file = append(file, "")
	file = append(file, "[[6]]")
	file = append(file, "[[2]]")

	packets := make([]d13Composite, 0)

	sum := 0
	for i := 0; i < len(file); i += 3 {
		first, second := file[i][1:len(file[i])-1], file[i+1][1:len(file[i+1])-1]
		c1, c2 := d13Composite{-1, make([]d13Composite, 0), nil}, d13Composite{-1, make([]d13Composite, 0), nil}

		fillComposite(&first, &c1)
		fillComposite(&second, &c2)

		if cmpComposite(c1, c2) == true {
			sum += i/3 + 1
		}

		packets = append(packets, c1, c2)
	}
	fmt.Println("Part 1:", sum)

	dividerOne := packets[len(packets)-2]
	dividerTwo := packets[len(packets)-1]

	sort.Slice(packets, func(i, j int) bool {
		return cmpComposite(packets[i], packets[j])
	})

	mulOne := 0
	mulTwo := 0
	for ii, i := range packets {
		if reflect.DeepEqual(i, dividerOne) {
			mulOne = ii + 1
		}
		if reflect.DeepEqual(i, dividerTwo) {
			mulTwo = ii + 1
		}

	}

	fmt.Println("Part 2:", mulOne*mulTwo)
}
