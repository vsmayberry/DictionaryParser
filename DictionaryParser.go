package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	runOpt := os.Args[2]
	InitAbbreviationMap()
	InitTonganWordMap()
	InitEnglishWordMap()
	if runOpt == "file" {
		fmt.Println("reading File")
	} else {
	http.HandleFunc("/parse", func(w http.ResponseWriter, req *http.Request) {
		body, _ := ioutil.ReadAll(req.Body)
		definitions, err := parsePage(string(body))
		if err != nil {
			io.WriteString(w, err.Error())
		}
		resp, err := json.Marshal(definitions)
			io.WriteString(w, string(resp))
		})
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}

func parsePage(dictionaryPage string) ([]Definition, error) {
	lines := strings.Split(dictionaryPage, "\n\n")
	return parsePageSlice(lines)
}

func parsePageSlice(dictionaryPageSlice []string) ([]Definition, error) {
	var definitions = make([]Definition, 50)
	defintionCount := 0
	for _, element := range dictionaryPageSlice {
		wordsInElement := strings.Split(element, " ")
		wordCountInElement := len(wordsInElement)
		if wordCountInElement == 1 {
			fmt.Println("\n\n\nFound header: " + element)
		} else if wordCountInElement == 2 {
			fmt.Println("\n\n\nFound second Header: " + element)
		} else {
			definition, err := isValidDefinition(element)
			if err == nil {
				fmt.Println("\n\n\nFound valid definition: " + definition.Word)
				definitions[defintionCount] = definition
				defintionCount++
			}
		}
	}
	definitions = definitions[0:defintionCount]
	return definitions, nil
}
