package main

import (
	"os"
	"strings"
    "fmt"
    "strconv"
)

func main() {
	input, _ := os.ReadFile("input")

	split := strings.Fields((string(input)))
    //fmt.Print(split)
	contain := 0
    overlap := 0
	for _, pair := range split {
        fmt.Print(pair, "\n")
		splitPair := strings.Split(pair, ",")
		elf1 := splitPair[0]
        fmt.Print(elf1, "\n")
		elf2 := splitPair[1]
        fmt.Print(elf2, "\n")

		if contains(elf1, elf2) {
            contain++
            //overlap++
        }
        if overlaps(elf1, elf2) {
            overlap++
        }
	}
    fmt.Print(contain, "\n")
    fmt.Print(overlap, "\n")
}

func contains(elf1 string, elf2 string) bool {
	elf1Boundries := strings.Split(elf1, "-") // index 0 is the left boundry
	elf2Boundries := strings.Split(elf2, "-") // index 1 is the right boundry
    fmt.Printf("Elf1 split: %v, Elf2 split: %v \n", elf1Boundries, elf2Boundries)
    e1Left, _ := strconv.Atoi(elf1Boundries[0])
    e1Right, _ := strconv.Atoi(elf1Boundries[1])
    fmt.Printf("Elf1 left boundry: %v, Elf1 right boundry %v \n", e1Left, e1Right)
    e2Left, _ := strconv.Atoi(elf2Boundries[0])
    e2Right, _ := strconv.Atoi(elf2Boundries[1])
    fmt.Printf("Elf2 left boundry: %v, Elf2 right boundry %v \n", e2Left, e2Right)
    
    // check if elf1 contains elf2
    if e1Left <= e2Left && e1Right >= e2Right {
        fmt.Printf("%v contains %v \n", elf1, elf2)
        return true
    }
    // check if elf2 contains elf1
    if e2Left <= e1Left && e2Right >= e1Right {
        fmt.Printf("%v contains %v \n\n", elf2, elf1)
        return true
    }
    return false
}


func overlaps(elf1 string, elf2 string) bool {
	elf1Boundries := strings.Split(elf1, "-") // index 0 is the left boundry
	elf2Boundries := strings.Split(elf2, "-") // index 1 is the right boundry
    fmt.Printf("Elf1 split: %v, Elf2 split: %v \n", elf1Boundries, elf2Boundries)
    e1Left, _ := strconv.Atoi(elf1Boundries[0])
    e1Right, _ := strconv.Atoi(elf1Boundries[1])
    fmt.Printf("Elf1 left boundry: %v, Elf1 right boundry %v \n", e1Left, e1Right)
    e2Left, _ := strconv.Atoi(elf2Boundries[0])
    e2Right, _ := strconv.Atoi(elf2Boundries[1])
    fmt.Printf("Elf2 left boundry: %v, Elf2 right boundry %v \n", e2Left, e2Right)
    
    if e1Left <= e2Left && e1Right >= e2Left {
        return true
    }
    if e2Left <= e1Left && e2Right >= e1Left {
        return true
    }
    if e1Left <= e2Right && e1Right >= e2Right {
        return true
    }
    if e2Left <= e1Right && e2Right >= e1Right {
        return true
    }

    return false
}
