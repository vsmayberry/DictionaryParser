package main

import "testing"

func TestSimpleEnglishWord(t *testing.T) {
	InitEnglishWordMap()
	if !IsEnglishWord("habitually") {
		t.Fatal("Failed to validate: habitually")
	}
}

func TestBritishSpellingEnglishWord(t *testing.T) {
	InitEnglishWordMap()
	if !IsEnglishWord("honourable") {
		t.Fatal("Failed to validate: honourable")
	}
}

func TestHyphenatedEnglishWord(t *testing.T) {
	InitEnglishWordMap()
	if !IsEnglishWord("rough-shod") {
		t.Fatal("Failed to validate: rough-shod")
	}
}
