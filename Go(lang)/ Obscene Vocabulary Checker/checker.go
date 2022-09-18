package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	fileName string
	wordList = make(map[string]struct{})
)

func main() {
	err := getFileName()
	if err != nil {
		log.Fatalf("Error at getFileName(): %s\n", err.Error())
	}

	err = readFile(fileName)
	if err != nil {
		log.Fatalf("Error at readFile(): %s\n", err.Error())
	}

	for {
		handleInput(getUserSentences())
	}
}

func getFileName() error {
	_, err := fmt.Scanln(&fileName)
	if err != nil {
		return err
	}
	return err
}

func readFile(name string) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordList[strings.ToLower(scanner.Text())] = struct{}{}
	}

	err = scanner.Err()
	if err != nil {
		return err
	}

	return nil
}

func getUserSentences() []string {
	var result []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	result = strings.Split(scanner.Text(), " ")
	return result
}

func handleInput(userSentence []string) {
	if len(userSentence) == 1 && userSentence[0] == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
	}
	writeOutput(userSentence)
}

func writeOutput(someWords []string) {
	fmt.Println(wordProcessing(someWords))
}

func wordProcessing(someWords []string) string {
	var b strings.Builder
	for index, word := range someWords {
		if needsCensorship(word) {
			b.WriteString(strings.Repeat("*", len(word)))
		} else {
			b.WriteString(word)
		}
		if index != len(someWords)-1 {
			b.WriteString(" ")
		}
	}
	return b.String()
}

func needsCensorship(someWord string) bool {
	_, ok := wordList[strings.ToLower(someWord)]
	return ok
}
