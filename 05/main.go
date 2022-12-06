package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var cargo = []Stack{
	{},
	{"N", "S", "D", "C", "V", "Q", "T"},
	{"M", "F", "V"},
	{"F", "Q", "W", "D", "P", "N", "H", "M"},
	{"D", "Q", "R", "T", "F"},
	{"R", "F", "M", "N", "Q", "H", "V", "B"},
	{"C", "F", "G", "N", "P", "W", "Q"},
	{"W", "F", "R", "L", "C", "T"},
	{"T", "Z", "N", "S"},
	{"M", "S", "D", "J", "R", "Q", "H", "N"},
}

type Stack []string

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Print() {
	printable := ""
	for _, element := range *s {
		printable += element
	}
	fmt.Println(printable)
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) PushMultiple(elements []string) {
	*s = append(*s, elements...)
}

func (s *Stack) PopMultiple(howMany int) ([]string, bool) {
    var elements []string
    for i:=0;i<howMany;i++{
        element, _ := (*s).Pop()
        elements = append(elements, element)
    }
    func(s []string) {
        sort.SliceStable(s, func(i, j int)  bool{
            return i>j
        })
    }(elements)
    return elements, true
}

func executeMove(cargo *[]Stack, move []int) *[]Stack {
	fmt.Println(move)
	(*cargo)[move[1]].Print()
	if move[0] == 1 {
		container, _ := (*cargo)[move[1]].Pop()
		(*cargo)[move[2]].Push(container)
	} else {
		containers, _ := (*cargo)[move[1]].PopMultiple(move[0])
		(*cargo)[move[2]].PushMultiple(containers)
	}
	return cargo
}

func printResult(cargo *[]Stack) {
	for _, stack := range *cargo {
		container, _ := stack.Pop()
		fmt.Println(container)
	}
}

func main() {
	input, _ := os.ReadFile("input")
	moves := strings.Split(strings.Split(string(input), "\n\n")[1], "\n")
	var cargoPostMoves *[]Stack
	for _, move := range moves {
		split := strings.Split(move, " ")
		howMany, _ := strconv.Atoi(split[1])
		from, _ := strconv.Atoi(split[3])
		to, _ := strconv.Atoi(split[5])
		extracted := []int{howMany, from, to}
		cargoPostMoves = executeMove(&cargo, extracted)
	}

	printResult(cargoPostMoves)

}
