package main

import "testing"

func TestOppAbbreviation(t *testing.T) {
	InitAbbreviationMap()
	if !IsAbbreviation("opp.") {
		t.Fatal("failed to find opp. abbreviation")
	}
}
