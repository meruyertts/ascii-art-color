package main

import (
	"ascii-art-color/checks"
	"ascii-art-color/splitprint"
	"fmt"
	"os"
	"strings"
)

func main() {
	color := "\033[97m"
	input := os.Args[1:]
	if len(input) < 2 || len(input) > 3 {
		fmt.Println("Usage: go run . [STRING] [OPTION]\nEX: go run . something --color=<color>")
		return
	}
	myStr := input[0]
	if !checks.IsASCII(myStr) {
		fmt.Println("non-ASCII character was entered")
		return
	}
	if len(myStr) == 0 {
		return
	}
	if checks.LineCounter() != nil {
		return
	}
	if len(input) >= 2 {
		c, err := checks.MyColor(strings.ToLower(input[1]))
		if c == "no match" {
			fmt.Println("Available colors are: red, green, yellow, blue, purple, cyan, gray, white, black and orange")
			return
		}
		if err != nil {
			fmt.Println("Usage: go run . [STRING] [OPTION]\nEX: go run . something --color=<color>")
			return
		}
		color = c
	}
	l := "[:]"
	if len(input) > 2 {
		l = input[2]
	}
	splitprint.PrintColoredWord(myStr, color, l)
}
