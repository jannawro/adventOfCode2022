package main

import (
	"fmt"
	"os"
	"strings"
)

const(
    k=14
    //k=4
)

func main() {
	input, _ := os.ReadFile("input")
	inputString := strings.TrimSpace(string(input))
	split := strings.Split(inputString, "")
    for i:=0;i < len(split)-k+1; i++ {
        tempSlice := split[i:i+k]
        fmt.Println(tempSlice)
        if len(tempSlice) == len(uniqueNonEmptyElementsOf(tempSlice)) {
            fmt.Printf("Found it: %v, after %v", tempSlice, i+k)
            break
        }
    }
}

func uniqueNonEmptyElementsOf(s []string) []string {
  unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}

	return us

}
