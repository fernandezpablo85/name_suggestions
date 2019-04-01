package main

import (
	"sort"
	"strings"
)

// Name is a valid name and it's frequency in the dictionary
type Name struct {
	word string
	freq int
}

type byFreq []Name

func (l byFreq) Len() int {
	return len(l)
}

func (l byFreq) Less(i, j int) bool {
	return l[i].freq > l[j].freq
}

func (l byFreq) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// Corrections returns the possible corrections ordered by frequency or nil if none.
func Corrections(s string) []string {
	alts := CorrectionsWithFrequency(s)
	ws := []string{}
	for _, a := range alts {
		ws = append(ws, a.word)
	}
	return ws
}

// CorrectionsWithFrequency returns the possible corrections and their frequency, ordered by frequency.
func CorrectionsWithFrequency(s string) []Name {
	lowerS := strings.ToLower(s)
	vars := Variations(lowerS)
	alts := []Name{}
	for _, v := range vars {
		if isOnDict(v) {
			alts = append(alts, Name{word: v, freq: Dict[v]})
		}
	}
	sort.Sort(byFreq(alts))
	return alts

}

func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}
	for index := range elements {
		if encountered[elements[index]] == false {
			// Record this element as an encountered element.
			encountered[elements[index]] = true
			// Append to result slice.
			result = append(result, elements[index])
		}
	}
	// Return the new slice.
	return result
}

// Variations returns all 1 edit variations for the given string.
func Variations(s string) []string {
	variations := []string{}
	pairs := pairs(s)
	variations = append(variations, deletions(pairs)...)
	variations = append(variations, transposes(pairs)...)
	variations = append(variations, replaces(pairs)...)
	variationsNoDuplicates := removeDuplicates(variations)
	return variationsNoDuplicates
}

type pair struct {
	first  string
	second string
}

func pairs(s string) []pair {
	ps := []pair{}
	for i := 0; i < len(s)+1; i++ {
		pair := pair{first: s[:i], second: s[i:]}
		ps = append(ps, pair)
	}
	return ps
}

func deletions(pairs []pair) []string {
	deletions := []string{}
	for _, p := range pairs {
		if len(p.second) == 0 {
			continue
		}
		del := p.first + p.second[1:]
		deletions = append(deletions, del)
	}
	return deletions
}

func transposes(pairs []pair) []string {
	ts := []string{}
	for _, p := range pairs {
		if len(p.second) <= 1 {
			continue
		}
		t := p.first + string(p.second[1]) + string(p.second[0]) + p.second[2:]
		ts = append(ts, t)
	}
	return ts
}

func replaces(pairs []pair) []string {
	letters := "abcdefghijklmnopqrstuvwxyz "
	rs := []string{}
	for _, p := range pairs {
		for _, l := range letters {
			i := p.first + string(l) + p.second
			rs = append(rs, i)
			if len(p.second) > 0 {
				r := p.first + string(l) + p.second[1:]
				rs = append(rs, r)
			}
		}
	}
	return rs
}
