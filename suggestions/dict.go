package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// Dict contains names and number of people with that name.
var Dict = map[string]int{}

func loadDict() {
	if len(Dict) > 0 {
		// already loaded
		return
	}
	file, err := os.Open(DictPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		name := parts[0]
		count, _ := strconv.Atoi(parts[1])
		if count > 80 {
			total = total + 1
			Dict[name] = count
		}

	}
	log.Printf("%d names loaded", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func isOnDict(name string) bool {
	_, ok := Dict[name]
	return ok
}

func init() {
	LoadDictFromRelativePathIfRunningTests()
}

// LoadDictFromRelativePathIfRunningTests changes DictPath to a relative, known path if the running command
// looks like a test run.
func LoadDictFromRelativePathIfRunningTests() {
	currentCommand := os.Args[0]
	testMode := strings.HasSuffix(currentCommand, ".test")
	if len(DictPath) <= 0 && testMode {
		DictPath = "../assets/dict.csv"
	}
}

func findSuggestionsFor(name string) ([]Name, bool) {
	loadDict()
	if isOnDict(name) {
		// name exists
		return nil, true
	}

	return CorrectionsWithFrequency(name), false
}
