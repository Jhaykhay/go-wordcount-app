package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	//Read contents of textfile
	content, err := ioutil.ReadFile("myfile.txt")
	if err != nil {
		panic(err)
	}

	//convert cotent into string
	text := string(content)

	//split text into words
	words := strings.Fields(text)

	//count words and spaces
	wordCount := len(words)
	spaceCounts := strings.Count(text, "")

	//findinf duplicate words and count their occurence
	duplicateMap := make(map[string]int)

	for _, word := range words {
		if _, ok := duplicateMap[word]; ok {
			duplicateMap[word]++
		} else {
			duplicateMap[word] = 1
		}
	}

	//print results
	fmt.Printf("Word count: %d\n", wordCount)
	fmt.Printf("Space count: %d\n", spaceCounts)
	fmt.Println("Duplicate words:")
	for word, count := range duplicateMap {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
}
