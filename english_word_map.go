package main

import (
	"bufio"
	"os"
	"strings"
)

//EnglishWordMap that holds all know valid english words
var EnglishWordMap map[string]int

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//InitEnglishWordMap initializes map of english words
func InitEnglishWordMap() {
	EnglishWordMap = make(map[string]int)
	loadAmericanWords()
	loadBritishWords()
}

func loadBritishWords() {
	f, err := os.Open("./british-english")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		word := scanner.Text()
		_, wordExists := EnglishWordMap[word]
		if !wordExists {
			EnglishWordMap[word] = 1
		}
	}
}

func loadAmericanWords() {
	f, err := os.Open("./words")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		word := scanner.Text()
		_, wordExists := EnglishWordMap[word]
		if !wordExists {
			EnglishWordMap[word] = 1
		}
	}
}

//IsEnglishWord returns whether the provided word is a known english word
func IsEnglishWord(possibleWord string) bool {
	possibleWord = strings.TrimPrefix(possibleWord, "(")
	possibleWord = strings.TrimSuffix(possibleWord, ")")
	possibleWord = strings.TrimSuffix(possibleWord, ",")
	possibleWord = strings.TrimSuffix(possibleWord, ".")
	possibleWord = strings.TrimSuffix(possibleWord, ":")
	possibleWord = strings.TrimSuffix(possibleWord, ";")
	possibleWord = strings.ToLower(possibleWord)
	_, exists := EnglishWordMap[possibleWord]
	if !exists && strings.Contains(possibleWord, "-") {
		rootWords := strings.Split(possibleWord, "-")
		exists = IsEnglishWord(rootWords[0]) && IsEnglishWord(rootWords[1])
	} else if !exists && strings.Contains(possibleWord, "\n") {
		rootWords := strings.Split(possibleWord, "\n")
		exists = IsEnglishWord(rootWords[0]) && IsEnglishWord(rootWords[1])
	} else if !exists && strings.Contains(possibleWord, "-\n") {
		rootWords := strings.Split(possibleWord, "-\n")
		exists = IsEnglishWord(rootWords[0]) && IsEnglishWord(rootWords[1])
	}
	return exists
}
