package main

import (
	"fmt"
	"strings"
)

func runLengthEncoding(input string ) string {
	count := 0
	var sb strings.Builder

	for i := 0; i < len(input); i++ {
		// first time	
		if i == 0 {
			count = 1
			continue
		}
		
		previous := input[i-1]
		current := input[i]

		// the same character was found
		if current == previous {
			count++
			continue
		}

		// a different char was found
		if count >  1 {
			sb.WriteString(fmt.Sprintf("%d", count))
		}
		sb.WriteByte(previous)
		count = 1
	}

	// Last char
	if count > 1 {
		sb.WriteString(fmt.Sprintf("%d", count))
	}
	sb.WriteByte(input[len(input)-1])

	return sb.String()

}

func main() {
	arr := []string{
		"aaabbcdddd",
		"BBBBBBBBBBBBNBBBBBBBBBBBBNNNBBBBBBBBBBBBBBBBBBBBBBBBNBBBBBBBBBBBBBB",
		"abcdefghijklmnopqrstuvwxyz",
	}

	for _, str := range arr {
		fmt.Println(str)
		fmt.Println(runLengthEncoding(str))
		println("----------------------------------")
	}
}