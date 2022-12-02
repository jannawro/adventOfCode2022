package main

import (
	"fmt"
	"os"
	"strings"
)

type Round struct {
	opponentsPlay string
	myPlay        string
}

type MyPlay struct {
	beats    string
	isBeatBy string
	draws    string
	points   int
}

type Guide map[string]MyPlay

type Tournament []Round

func parseInput(path string) string {
	input, _ := os.ReadFile(path)
	return string(input)
}

func makeTournament(input string) Tournament {
	var tournament []Round
	splitInput := strings.Split(input, "\n")
	for _, round := range splitInput {
		roundPlays := strings.Split(round, " ")
		tournament = append(tournament, Round{
			opponentsPlay: roundPlays[0],
			myPlay:        roundPlays[1],
		})
	}
	return tournament
}

func makeGuide() Guide {
	guide := map[string]MyPlay{
		"X": {
			beats:    "C",
			isBeatBy: "B",
			draws:    "A",
			points:   1,
		},
		"Y": {
			beats:    "A",
			isBeatBy: "C",
			draws:    "B",
			points:   2,
		},
		"Z": {
			beats:    "B",
			isBeatBy: "A",
			draws:    "C",
			points:   3,
		},
	}

	return guide
}

func (t Tournament) play(guide Guide) {
	pointsTotal := 0
	for _, round := range t {
		pointsTotal += round.play(guide)
	}
	fmt.Println(pointsTotal)
}

func (t Tournament) playRigged() {
	pointsTotal := 0
	for _, round := range t {
		pointsTotal += round.playRigged()
	}
	fmt.Println(pointsTotal)
}

func (r Round) play(guide Guide) int {
	return guide.compareAgainst(r)
}

func (r Round) playRigged() int {
	points := 0
	switch r.myPlay {
	case "X": // need to lose
		points += 0
		switch r.opponentsPlay {
		case "A":
			points += 3
		case "B":
			points += 1
		case "C":
			points += 2
		}
	case "Y": // need to draw
		points += 3
		switch r.opponentsPlay {
		case "A":
			points += 1
		case "B":
			points += 2
		case "C":
			points += 3
		}
	case "Z": // need to win
		points += 6
		switch r.opponentsPlay {
		case "A":
			points += 2
		case "B":
			points += 3
		case "C":
			points += 1
		}
	}
	return points
}

func (g Guide) compareAgainst(round Round) int {
	points := 0
	points += g[round.myPlay].points
	if round.opponentsPlay == g[round.myPlay].beats {
		points += 6
	} else if round.opponentsPlay == g[round.myPlay].isBeatBy {
		points += 0
	} else if round.opponentsPlay == g[round.myPlay].draws {
		points += 3
	}
	return points
}

func main() {
	input := parseInput("input")
	tournament := makeTournament(input)
	tournament.play(makeGuide())
	tournament.playRigged()
}
