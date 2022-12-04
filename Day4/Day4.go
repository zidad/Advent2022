package main

import (
	advent "Advent2022"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := make(chan string, 2)
	go advent.ReadInput(lines)

	var count int
	var countOverlaps int
	for line := range lines {
		ranges := strings.Split(line, ",")
		range1, _ := sliceAtoi(strings.Split(ranges[0], "-"))
		range2, _ := sliceAtoi(strings.Split(ranges[1], "-"))

		containsFully := contains(range1, range2) || contains(range2, range1)
		overlaps := overlaps(range1, range2)

		if containsFully {
			fmt.Printf("%s range1 containsFully range2 or vice versa \n", line)
			count++
		}
		if overlaps {
			fmt.Printf("%s range1 partially overlaps range2 or vice versa \n", line)
			countOverlaps++
		}
	}
	fmt.Printf("contains %d overlaps %d", count, countOverlaps)
}

func contains(range1 []int, range2 []int) bool {
	return range1[0] <= range2[0] && range1[1] >= range2[1]
}

func overlaps(range1 []int, range2 []int) bool {
	return between(range1[0], range2[0], range2[1]) || between(range2[0], range1[0], range1[1])
}

func between(i, min, max int) bool {
	return (i >= min) && (i <= max)
}

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}
