package main

import (
	"testing"
)

func TestSimpleWordDef(t *testing.T) {
	InitTonganWordMap()
	InitEnglishWordMap()
	_, err := isValidDefinition("anga-faitotonu, habitually honest or honourable in one's dealings.")
	if err != nil {
		t.Fatal("failed to parse simple word class definition example")
	}
}

func TestSimpleWordDef2(t *testing.T) {
	InitTonganWordMap()
	InitEnglishWordMap()
	definition, err := isValidDefinition("anga-fakafefine, habitually behaving like a woman or a girl, womanish, girlish.")
	if err != nil {
		t.Fatal("failed to parse simple word class definition example")
	}
	if definition.Word != "anga-fakafefine" {
		t.Fatal("Failed to build definition")
	}
}

func TestSimpleWordDefWithAbbreviation(t *testing.T) { //test fails anga-fufi shoud be anga-fufuu
	InitTonganWordMap()
	InitEnglishWordMap()
	InitAbbreviationMap()
	_, err := isValidDefinition("anga-fakahaha, demonstrative: opp. anga-fufi.")
	if err == nil {
		t.Fatal("failed to parse simple word class definition with abbreviation example")
	}
}

func TestTonganPhraseMatching(t *testing.T) {
	InitTonganWordMap()
	InitEnglishWordMap()
	InitAbbreviationMap()
	_, err := isValidDefinition(`fai tu‘u, v.i., to do standing or while
	standing, — see tu‘u?; (of a sickness)
	to be so serious that the people in the
	house are standing or on the move all
	the time with anxiety or in the en-
	deavour to do something for the sick
	person. ‘Oku fai tu‘u hono mahaki.
	(Said when death seems imminent or
	certain.)`)
	if err != nil {
		t.Fatal("failed to parse definition with tongan phrase")
	}
}
