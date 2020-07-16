package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//TonganWordMap that holds all know valid tongan words
var TonganWordMap map[string]int

//InitTonganWordMap initializes map of tongan words
func InitTonganWordMap() {
	TonganWordMap = make(map[string]int)
	f, err := os.Open("./tonganwordlist.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		word := scanner.Text()
		_, wordExists := TonganWordMap[word]
		if !wordExists {
			TonganWordMap[word] = 1
		}
	}

}

func addNewCompoundWord(word string) {
	f, err := os.OpenFile("./tonganwordlist.txt", os.O_WRONLY|os.O_APPEND, 0644)
	check(err)
	defer f.Close()

	_, err = f.WriteString(word + "\n")
	check(err)
}

//IsTonganPhrase returns if the given string is a set of valid tongan words
func IsTonganPhrase(possiblePhrase string) bool {
	wordsInPhrase := strings.Split(possiblePhrase, " ")
	for _, word := range wordsInPhrase {
		if !IsTonganWord(word) {
			return false
		}
	}
	return true
}

//IsTonganWord returns whether the provided word is a known tongan word
func IsTonganWord(possibleWord string) bool {
	possibleWord = strings.TrimPrefix(possibleWord, "(")
	possibleWord = strings.TrimSuffix(possibleWord, ")")
	possibleWord = strings.TrimSuffix(possibleWord, ",")
	possibleWord = strings.TrimSuffix(possibleWord, ".")
	possibleWord = strings.TrimSuffix(possibleWord, ":")
	possibleWord = strings.TrimSuffix(possibleWord, ";")
	possibleWord = strings.ToLower(possibleWord)
	_, exists := TonganWordMap[possibleWord]
	if !exists {
		exists = isCompoundWord(possibleWord)
		if exists {
			addNewCompoundWord(possibleWord)
		}
	}
	return exists
}

func isCompoundWord(possibleWord string) bool {
	return isAngaWord(possibleWord) || isFakaWord(possibleWord) || isHyphenatedWord(possibleWord)
}

func isHyphenatedWord(possibleWord string) bool {
	if strings.Contains(possibleWord, "-") {
		for _, partialWord := range strings.Split(possibleWord, "-") {
			fmt.Println("partial word:" + partialWord + ":")
			if IsTonganWord(partialWord) {
				return true
			}
		}
	}
	return true
}

func isFakaWord(possibleWord string) bool {
	return strings.HasPrefix(possibleWord, "faka") &&
		IsTonganWord(strings.TrimPrefix(possibleWord, "faka"))
}

func isAngaWord(possibleWord string) bool {
	return strings.HasPrefix(possibleWord, "anga-") &&
		IsTonganWord(strings.TrimPrefix(possibleWord, "anga-"))
}
