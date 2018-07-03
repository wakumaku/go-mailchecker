package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 5 {
		printHelp(nil)
		os.Exit(1)
	}

	fileIn := os.Args[1]
	fileOut := os.Args[2]
	fileJSON := os.Args[3]
	keyword := os.Args[4]

	template, err := ioutil.ReadFile(fileIn)
	if err != nil {
		printHelp(err)
		os.Exit(1)
	}

	jsonContent, err := ioutil.ReadFile(fileJSON)
	if err != nil {
		printHelp(err)
		os.Exit(1)
	}

	var domainMap [][]string
	if err := json.Unmarshal(jsonContent, &domainMap); err != nil {
		printHelp(err)
		os.Exit(1)
	}

	var domainList string
	for _, segment := range domainMap {
		for _, domain := range segment {
			domainList += fmt.Sprintf("\"%s\",", domain)
		}
	}

	contentFile := strings.Replace(string(template), keyword, domainList, -1)

	if err := ioutil.WriteFile(fileOut, []byte(contentFile), 0644); err != nil {
		printHelp(err)
		os.Exit(1)
	}

	// no news, good news
}

func printHelp(err error) {
	fmt.Println("Error:", err)
	fmt.Println("Wrong arguments!")
}
