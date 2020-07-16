package main

import (
	"bufio"
	"os"
	"strings"
)

//AbbreviationMap that holds all know valid tongan words
var AbbreviationMap map[string]int

//InitAbbreviationMap initializes map of tongan words
func InitAbbreviationMap() {
	AbbreviationMap = make(map[string]int)
	f, err := os.Open("./abbreviations.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		word := scanner.Text()
		_, wordExists := AbbreviationMap[word]
		if !wordExists {
			AbbreviationMap[word] = 1
		}
	}
}

//IsAbbreviation returns whether the provided word is a known tongan word
func IsAbbreviation(possibleWord string) bool {
	possibleWord = strings.TrimPrefix(possibleWord, "(")
	possibleWord = strings.TrimSuffix(possibleWord, ")")
	possibleWord = strings.TrimSuffix(possibleWord, ",")
	possibleWord = strings.TrimSuffix(possibleWord, ":")
	possibleWord = strings.TrimSuffix(possibleWord, ";")
	possibleWord = strings.ToLower(possibleWord)
	_, exists := AbbreviationMap[possibleWord]
	return exists
}
