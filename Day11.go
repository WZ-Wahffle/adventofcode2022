package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items        []int64
	opSymbol     rune
	opValue      int64
	testValue    int64
	ifTrue       int
	ifFalse      int
	inspectCount int
}

func main() {
	dat, _ := os.ReadFile("input/Day11Input.txt")
	file := strings.Split(string(dat), "\n")

	fmt.Print("Run part 1 or part 2? ")
	part := 0
	_, _ = fmt.Scanf("%d", &part)
	if part != 1 && part != 2 {
		os.Exit(1)
	}

	monkeys := make([]monkey, 0)
	lastIdx := -1

	for i := 0; i < len(file); i++ {
		monkeys = append(monkeys, monkey{})
		lastIdx++
		i++

		opsBuffer := file[i][18:]
		ops := strings.Split(opsBuffer, ", ")
		for _, j := range ops {
			temp, _ := strconv.ParseInt(j, 10, 64)

			monkeys[lastIdx].items = append(monkeys[lastIdx].items, temp)
		}
		i++

		if file[i][25] == 'o' {
			_, _ = fmt.Sscanf(file[i], "  Operation: new = old %c old", &monkeys[lastIdx].opSymbol)
			monkeys[lastIdx].opValue = -1
		} else {
			_, _ = fmt.Sscanf(file[i], "  Operation: new = old %c %d", &monkeys[lastIdx].opSymbol, &monkeys[lastIdx].opValue)
		}
		i++

		_, _ = fmt.Sscanf(file[i], "  Test: divisible by %d", &monkeys[lastIdx].testValue)
		i++

		_, _ = fmt.Sscanf(file[i], "    If true: throw to monkey %d", &monkeys[lastIdx].ifTrue)
		i++

		_, _ = fmt.Sscanf(file[i], "    If false: throw to monkey %d", &monkeys[lastIdx].ifFalse)
		i++
	}

	// getting all values for 10000 iterations is not constrained by time, but by integer size
	// as such, all worry values are mod'd by the LCM of all test values in order to keep them in check
	//
	// let it be known that this was only the second option after attempting to use a slightly dysfunctional
	// int128 library, which would've probably worked as well
	masterMod := int64(1)

	for _, i := range monkeys {
		masterMod *= i.testValue
	}

	iterations := 20
	if part == 2 {
		iterations = 10000
	}
	for range iterations {
		for i := range monkeys {
			for range monkeys[i].items {
				switch monkeys[i].opSymbol {
				case '+':
					if monkeys[i].opValue == -1 {
						monkeys[i].items[0] += monkeys[i].items[0]
					} else {
						monkeys[i].items[0] += monkeys[i].opValue
					}
					break
				case '*':
					if monkeys[i].opValue == -1 {
						monkeys[i].items[0] *= monkeys[i].items[0]
					} else {
						monkeys[i].items[0] *= monkeys[i].opValue
					}
					break
				}
				if part == 1 {
					monkeys[i].items[0] /= 3
				} else {
					monkeys[i].items[0] %= masterMod
				}

				if monkeys[i].items[0]%monkeys[i].testValue == 0 {
					monkeys[monkeys[i].ifTrue].items = append(monkeys[monkeys[i].ifTrue].items, monkeys[i].items[0])
				} else {
					monkeys[monkeys[i].ifFalse].items = append(monkeys[monkeys[i].ifFalse].items, monkeys[i].items[0])
				}
				monkeys[i].items = monkeys[i].items[1:]
				monkeys[i].inspectCount++
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		if monkeys[i].inspectCount > monkeys[j].inspectCount {
			return true
		}
		return false
	})

	fmt.Printf("Part %d: %d", part, monkeys[0].inspectCount*monkeys[1].inspectCount)

}
