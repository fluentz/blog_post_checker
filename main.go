package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/mgutz/ansi"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Please provide the path of the file to check.")
	}

	path := os.Args[1]

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Unable to open file at given path.")
	}

	text := string(content)
	text = HighlightHomphones(text)
	text = HighlightApostrophes(text)
	fmt.Print(text)
}

func HighlightApostrophes(s string) string {
	chars := strings.Split(s, "")
	yellow := ansi.ColorCode("yellow")
	reset := ansi.ColorCode("reset")

	formatted := ""

	for _, char := range chars {
		if char == "'" {
			formatted += fmt.Sprint(yellow, char, reset)
		} else {
			formatted += fmt.Sprint(char)
		}
	}

	return formatted
}

func HighlightHomphones(s string) string {
	homophoneDict := buildMap()
	words := strings.Split(s, " ")
	yellow := ansi.ColorCode("yellow")
	reset := ansi.ColorCode("reset")

	formatted := ""

	for idx, word := range words {
		if isHomophone(homophoneDict, word) {
			formatted += fmt.Sprint(yellow, formatWord(idx, word), reset)
		} else {
			formatted += fmt.Sprint(formatWord(idx, word))
		}
	}

	return formatted
}

func formatWord(idx int, w string) string {
	if idx == 0 {
		return w
	} else {
		return " " + w
	}
}

func buildMap() map[string]bool {
	var m map[string]bool
	m = make(map[string]bool)

	words := strings.Split(homophones, "\n")
	for _, word := range words {
		m[word] = true
	}
	return m
}

func isHomophone(dict map[string]bool, s string) bool {
	return dict[strings.ToLower(s)]
}
