package splitprint

import (
	"ascii-art-color/checks"
	"ascii-art-color/read"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func PrintColoredWord(myStr, c, l string) {
	third := l
	re := regexp.MustCompile(`\\n`)
	newStr := re.Split(myStr, -1)
	myStr = strings.ReplaceAll(myStr, "\\n", "")
	indexstr := 0
	for i := 0; i < len(newStr); i++ {
		if len(newStr[i]) > 0 {
			if third == myStr {
				l = newStr[i]
			}
			toColor, err := checks.Index(myStr, l)
			if err != nil {
				log.Println(err)
				return
			}
			printWord(newStr[i], c, toColor, indexstr)
			indexstr += len(newStr[i])
		}
		if newStr[i] == "" {
			fmt.Println("")
		}
	}
}

func printWord(s string, c string, l []int, indexstr int) {
	myarray := [8]string{}
	myline := ""
	color := "\033[97m"
	for _, char := range s {
		for _, i := range l {
			if indexstr == i {
				color = c
				break
			}
		}
		for line := 2; line <= 9; line++ {
			myrune := int(char)
			for i := ' '; i <= '~'; i++ {
				j := (int(i) - ' ') * 9
				if myrune == int(i) {
					firstline, err := read.ReadExactLine("standard.txt", line+j)
					if err != nil {
						log.Print(err)
						return
					}
					myline += firstline
				}
			}
		}
		temp := strings.Split(myline, "\n")
		for index, s := range temp[:len(temp)-1] {
			myarray[index] += color + s + "\033[0m"
		}
		myline = ""
		color = "\033[97m"
		indexstr++
	}
	for _, s := range myarray {
		fmt.Println(s)
	}
}
