package main

import (
	"testing"
)

func BenchmarkSuggestions10(b *testing.B) { // about 160µs
	for n := 0; n < b.N; n++ {
		findSuggestionsFor("juan pablo")
	}
}

func TestSuggestion10Exists(t *testing.T) {
	name := "juan pablo"
	sugs, exists := findSuggestionsFor(name)
	if !exists {
		t.Errorf("dictionary should contain '%s'", name)
	}

	if len(sugs) > 0 {
		t.Errorf("there should be no suggestions for an existing name like '%s'", name)
	}

}

func TestSuggestion10NotExists(t *testing.T) {
	name := "juan pable"
	sugs, exists := findSuggestionsFor(name)
	if exists {
		t.Errorf("dictionary should not contain '%s'", name)
	}

	if len(sugs) <= 0 {
		t.Errorf("there should be suggestions for a typo like '%s'", name)
	}

	mostCommonSuggestion := "juan pablo"
	if !containsName(sugs, mostCommonSuggestion) {
		t.Errorf("suggestions must contain '%s', the most common one", mostCommonSuggestion)
	}
}

func containsName(names []Name, toFind string) bool {
	for _, n := range names {
		if n.word == toFind {
			return true
		}
	}
	return false
}
