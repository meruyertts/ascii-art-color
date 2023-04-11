package checks

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func IsASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}
func LineCounter() error {
	file, err := os.Open("standard.txt")
	if err != nil {
		return err
	}
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	if lineCount < 855 {
		return errors.New("file does not contain all characters")
	}
	return nil
}

func MyColor(myColorStr string) (string, error) {
	var err error
	output := "no match"
	re := regexp.MustCompile(`^--color=`)
	match := re.Split(myColorStr, -1)
	if len(match) == 1 || len(match[1]) == 0 {
		return "", errors.New("no color")
	}
	switch match[1] {
	case "red":
		output = "\033[31m"
	case "green":
		output = "\033[32m"
	case "yellow":
		output = "\033[33m"
	case "blue":
		output = "\033[34m"
	case "purple":
		output = "\033[35m"
	case "cyan":
		output = "\033[36m"
	case "gray":
		output = "\033[37m"
	case "white":
		output = "\033[97m"
	case "black":
		output = "\033[30m"
	case "orange":
		output = "\033[38;2;255;165;0m"
	}
	return output, err
}

func Index(first, third string) ([]int, error) {
	var errs error
	var indexValues []int
	re := regexp.MustCompile(`\[(.*?)\]`)
	indices := re.FindStringSubmatch(third)
	if len(indices) == 0 {
		fmt.Println("Usage: go run . [STRING] [OPTION] [INDEX]\nEX: go run . something --color=<color> [NUMBER1,NUMBER2] OR [NUMBER1:NUMBER2]")
		fmt.Println("EX: go run . \"hello, world\" --color=cyan \"[1:4]\"")
		return indexValues, errors.New("not valid input")
	}
	indexNum := indices[1]
	if len(indexNum) == 1 && indexNum != ":" {
		start, err := strconv.Atoi(indexNum)
		if err != nil {
			log.Println(err)
			return indexValues, err
		}
		indexValues = append(indexValues, start)
	} else if strings.Contains(indexNum, ":") {
		nums := strings.Split(indexNum, ":")
		if nums[0] == "" && nums[1] == "" {
			nums[0] = "0"
			nums[1] = strconv.Itoa(len(first) - 1)
		} else if nums[0] == "" {
			nums[0] = "0"
		} else if nums[1] == "" {
			nums[1] = strconv.Itoa(len(first) - 1)
		}
		start, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Println(err)
			return indexValues, err
		}
		end, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Println(err)
			return indexValues, err
		}
		if start < 0 || end < 0 {
			return indexValues, errors.New("negative number is entered")
		}
		if end < start {
			return indexValues, errors.New("invalid input")
		}
		for i := start; i <= end; i++ {
			indexValues = append(indexValues, i)
		}
	} else if strings.Contains(indexNum, ",") {
		nums := strings.Split(indexNum, ",")
		for _, nb := range nums {
			single, err := strconv.Atoi(strings.TrimSpace(nb))
			if err != nil {
				log.Println(err)
				return indexValues, err
			}
			indexValues = append(indexValues, single)
		}
	} else {
		fmt.Println("Usage: go run . [STRING] [OPTION] [INDEX]\nEX: go run . something --color=<color> [NUMBER1,NUMBER2] OR [NUMBER1:NUMBER2]")
		return indexValues, errors.New("not valid input")
	}
	return indexValues, errs
}
