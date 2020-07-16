package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func isValidDefinition(definitionString string) (Definition, error) {
	definitionString = removeInvalidCharacters(definitionString)
	definition, err := isWordDef(definitionString)
	return definition, err
}

func isWordDef(definition string) (Definition, error) {
	definitionParts := strings.Split(definition, ",")
	isTonganWord := IsTonganWord(definitionParts[0])
	isTonganPhrase := IsTonganPhrase(definitionParts[0])
	definitionBody, err := isDefinitionBody(definitionParts[1:])
	if (isTonganWord || isTonganPhrase) && err == nil {
		return Definition{definitionParts[0], definitionBody}, nil
	}
	return Definition{}, errors.New("Invalid Defintion")
}

func removeInvalidCharacters(originalString string) string {
	wordBreak := regexp.MustCompile(`\s*-\n\s*`)

	originalString = wordBreak.ReplaceAllString(originalString, "")
	originalString = strings.ReplaceAll(originalString, "—", " ")
	originalString = strings.ReplaceAll(originalString, "‘", "'")
	originalString = strings.ReplaceAll(originalString, "?", " ")
	originalString = strings.ReplaceAll(originalString, "‘", "'")
	originalString = strings.ReplaceAll(originalString, ";", ",")
	originalString = strings.ReplaceAll(originalString, "\n", "\n ")
	return originalString
}

func isDefinitionBody(possiblePartialDefinitions []string) ([]string, error) {
	var definitionArray []string = make([]string, 25)
	partialDefinitionCount := 0
	for _, partialDefinition := range possiblePartialDefinitions {
		partial, err := handlePartialDefinition(partialDefinition)
		if err != nil {
			return nil, errors.New("Invalid Word in defition: " + partial)
		}
		definitionArray[partialDefinitionCount] = partial
		partialDefinitionCount++
	}
	definitionArray = definitionArray[0:partialDefinitionCount]
	return definitionArray, nil
}
func handlePartialDefinition(possibleDefinitionBody string) (string, error) {
	possibleDefinitionBody = strings.ReplaceAll(possibleDefinitionBody, "\n", " ")
	space := regexp.MustCompile(`\s+`)
	possibleDefinitionBody = string(space.ReplaceAllString(possibleDefinitionBody, " "))
	for _, word := range strings.Split(possibleDefinitionBody, " ") {
		if word != "" && !IsEnglishWord(word) && !IsTonganWord(word) && !IsAbbreviation(word) {
			fmt.Println("Found invalid word: " + word)
			return word, errors.New("invalid word found in partial definition")
		}
	}
	return possibleDefinitionBody, nil
}
