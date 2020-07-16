package main

import "testing"

func TestSimpleTonganWord(t *testing.T) {
	InitTonganWordMap()
	if !IsTonganWord("ako") {
		t.Fatal("Failed to find word: ako")
	}
}

func TestFakafiemoheaTonganWord(t *testing.T) {
	InitTonganWordMap()
	if !IsTonganWord("fakafiemohea") {
		t.Fatal("Failed to find word: fakafiemohea")
	}
}

func TestHyphenatedTonganWord(t *testing.T) {
	InitTonganWordMap()
	if !IsTonganWord("fai-me‘a-lava") {
		t.Fatal("Failed to recognize hyphenated word")
	}
}

func TestAngafufiTonganWord(t *testing.T) { //anga-fufi is not a word anga-fufuu is
	InitTonganWordMap()
	if IsTonganWord("anga-fufi.") {
		t.Fatal("Failed to find word: anga-fufi.")
	}
}

func TestAngafakafefineTonganWord(t *testing.T) {
	InitTonganWordMap()
	if !IsTonganWord("anga-fakafefine") {
		t.Fatal("Failed to find word: anga-fakafefine")
	}
}

func TestAngafaitotonuTonganWord(t *testing.T) {
	InitTonganWordMap()
	if !IsTonganWord("anga-faitotonu") {
		t.Fatal("Failed to find word: anga-faitotonu")
	}
}

func TestFaituu(t *testing.T) {
	InitTonganWordMap()
	if !IsTonganPhrase("fai tu'u") {
		t.Fatal("Failed to find word: fai tu'u")
	}
}

func TestInvalidPhrase(t *testing.T) {
	InitTonganWordMap()
	if IsTonganPhrase("fakatonga bread") {
		t.Fatal("Failed to invalidate: fakatonga bread")
	}
}
