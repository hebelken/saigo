package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	// Read in file
	data, err := ioutil.ReadFile("7oldsamr.txt")
	if err != nil {
		panic(err)
	}

	// filter out punctuation
	replacer := strings.NewReplacer(",", "", ".", "", ";", "", "'", "", "!", "", "?", "", ":", "", "\"", "")
	parsedData := replacer.Replace(string(data))

	// split data from file by word
	words := strings.Fields(string(parsedData))

	// count all words as lowercase
	dict := map[string]int{}
	for _, word := range words {
		dict[strings.ToLower(word)] += 1
	}

	// reverse keys and values { 10: ["face", "cool", "gg"], 4: ["mm"], ... }
	reverseDict := map[int][]string{}
	counts := []int{}
	for word, count := range dict {
		reverseDict[count] = append(reverseDict[count], word)
	}

	// Grab all of the keys from the reversed dict to sort
	for count := range reverseDict {
		counts = append(counts, count)
	}

	// Sort
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	// Output
	for _, count := range counts {
		for _, word := range reverseDict[count] {
			fmt.Printf("%s %d\n", word, count)
		}
	}
}
