package main

import (
	advent "Advent2022"
	"fmt"
)

var names = map[uint8]string{'A': "Rock", 'B': "Paper", 'C': "Scissors"}
var beats = map[uint8]uint8{'A': 'C', 'B': 'A', 'C': 'B'}
var loses = map[uint8]uint8{'C': 'A', 'A': 'B', 'B': 'C'}
var points = map[uint8]int{'A': 1, 'B': 2, 'C': 3}

func main() {
	lines := make(chan string, 2)
	go advent.ReadInput(lines)

	opponentScoreTotal := 0
	myScoreTotal := 0
	round2 := true

	for line := range lines {
		opponent := line[0]
		myChoice := line[2] - 23

		if round2 {
			myChoice = DetermineMyChoice(line[2], opponent)
		}

		outcome := determineOutcome(opponent, myChoice)
		myScore, opponentScore := determineScore(outcome, myChoice, opponent)
		myScoreTotal += myScore
		opponentScoreTotal += opponentScore
		fmt.Printf("%s, opponent: %s %d me: %s %d outcome: %s points: %d %d \n", line, names[opponent], opponent, names[myChoice], myChoice, outcome, opponentScore, myScore)

	}
	fmt.Println("---")
	fmt.Printf("totals me: %d opponent: %d\n", myScoreTotal, opponentScoreTotal)
}

func determineScore(outcome string, me uint8, opponent uint8) (int, int) {
	switch outcome {
	case "draw":
		return points[me] + 3, points[opponent] + 3
	case "win":
		return points[me] + 6, points[opponent] + 0
	default:
		return points[me], points[opponent] + 6
	}
}

func determineOutcome(opponent uint8, me uint8) string {
	if opponent == me {
		return "draw"
	} else if beats[me] == opponent {
		return "win"
	} else {
		return "lose"
	}
}

func DetermineMyChoice(expectedOutcome uint8, opponent uint8) uint8 {
	switch expectedOutcome {
	case 'Y':
		return opponent
	case 'Z':
		return loses[opponent]
	default:
		return beats[opponent]
	}
}
