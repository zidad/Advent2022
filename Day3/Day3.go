package main

import (
	advent "Advent2022"
	"fmt"
)

func main() {
	lines := make(chan string, 2)
	go advent.ReadInput(lines)

	var totalScore int
	var totalBadgeScore int
	var lineNr int
	var last3lines []string
	for line := range lines {
		lineNr++
		last3lines = append(last3lines, line)
		compartmentA := line[0 : len(line)/2]
		compartmentB := line[len(line)/2:]

		intersected := Intersect([]uint8(compartmentA), []uint8(compartmentB))
		score := getScoreForItem(intersected)

		if lineNr%3 == 0 {
			badge := Intersect(Intersect([]uint8(last3lines[0]), []uint8(last3lines[1])), []uint8(last3lines[2]))
			badgeScore := getScoreForItem(badge)
			totalBadgeScore += badgeScore
			fmt.Printf("badge: %s %d \n", string(badge[0]), badgeScore)
			last3lines = []string{}
		}

		totalScore += score
		fmt.Printf("A %s B%s intersection: %s %d \n", compartmentA, compartmentB, string(intersected[0]), score)

	}
	fmt.Println("---")
	fmt.Printf("totalScored: %d bagdeScore %d \n", totalScore, totalBadgeScore)
}

func getScoreForItem(intersected []uint8) int {
	score := int(intersected[0])
	if intersected[0] >= 'a' {
		score -= 96
	} else {
		score -= 38
	}
	return score
}

func Intersect[T comparable](a []T, b []T) []T {
	set := make([]T, 0)

	for _, v := range a {
		if containsGeneric(b, v) {
			set = append(set, v)
		}
	}

	return set
}

func containsGeneric[T comparable](b []T, e T) bool {
	for _, v := range b {
		if v == e {
			return true
		}
	}
	return false
}
