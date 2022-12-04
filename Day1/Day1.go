package main

import (
	advent "Advent2022"
	"fmt"
	"log"
	"sort"
	"strconv"
)

func main() {
	lines := make(chan string, 2)
	go advent.ReadInput(lines)

	currentElf := 1
	runningTotalForCurrentElf := 0
	elfWithMostCalories := 0
	maxTotalCalories := 0
	var x []int

	for line := range lines {
		if line == "" {
			currentElf = currentElf + 1
			if runningTotalForCurrentElf > maxTotalCalories {
				elfWithMostCalories = currentElf
				maxTotalCalories = runningTotalForCurrentElf
			}
			fmt.Printf("elf %d totals %d\n", currentElf, runningTotalForCurrentElf)
			x = append(x, runningTotalForCurrentElf)
			runningTotalForCurrentElf = 0
		} else {
			value, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			runningTotalForCurrentElf = runningTotalForCurrentElf + value
		}
	}

	sort.Sort(sort.IntSlice(x))
	fmt.Println(x)

	y := x[len(x)-3:]
	fmt.Println(y)

	fmt.Println(sum(y))
	fmt.Println("---")
	fmt.Printf("elf %d has the highest amount of calories with %d\n", elfWithMostCalories, maxTotalCalories)
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
