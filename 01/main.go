package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := parseInput("input")
	elves := unmarshalToElves(input)
	findElfWithMostCalories(elves)
	findTotalOfTopThreeElves(elves)
}

type elf struct {
	position int
	calories []int
}

func (e elf) totalCarried() int {
	total := 0
	for _, pack := range e.calories {
		total += pack
	}
	return total
}

func parseInput(path string) string {
	input, _ := os.ReadFile(path)
	return string(input)
}

func unmarshalToElves(input string) []elf {
	var elves []elf
	splitInput := strings.Split(input, "\n\n")
	for i, calories := range splitInput {
		caloriesSplit := strings.Split(calories, "\n")
		var caloriesAsInt []int
		for _, calories := range caloriesSplit {
			integer, _ := strconv.Atoi(calories)
			caloriesAsInt = append(caloriesAsInt, integer)
		}
		elves = append(elves, elf{
			position: i,
			calories: caloriesAsInt,
		})
	}
	return elves
}

func findElfWithMostCalories(elves []elf) {
	var totals []int
	for _, elf := range elves {
		totals = append(totals, elf.totalCarried())
	}

	max := totals[0]
	for i := 1; i < len(totals); i++ {
		if max < totals[i] {
			max = totals[i]
		}
	}
	fmt.Println(max)
}

func findTotalOfTopThreeElves(elves []elf) {
	var totals []int
	for _, elf := range elves {
		totals = append(totals, elf.totalCarried())
	}

	top1 := -1
	top2 := -1
	top3 := -1
	for i := 0; i < len(totals); i++ {
		if totals[i] > top1 {
			top3 = top2
			top2 = top1
			top1 = totals[i]
		} else if totals[i] > top2 && totals[i] != top1 {
			top3 = top2
			top2 = totals[i]
		} else if totals[i] > top3 && totals[i] != top2 {
			top3 = totals[i]
		}
	}
	fmt.Println(top1 + top2 + top3)
}
