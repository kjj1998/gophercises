package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilePtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limitPtr := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilePtr)
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	correct := quiz(&data, *limitPtr)

	fmt.Printf("Number of questions answered correctly: %d\n", correct)
	fmt.Printf("Total number of questions: %d\n", len(data))
}

func quiz(questionsAndAnswers *[][]string, limit int) (correct int) {
	fmt.Printf("Press Enter to start the quiz, you have %d seconds", limit)
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')

	correct = 0
	scanner := bufio.NewScanner(os.Stdin)
	answerCh := make(chan string)
	end := time.After(time.Duration(limit) * time.Second)

	for index, row := range *questionsAndAnswers {
		fmt.Printf("Problem #%d: %s = ", index+1, row[0])

		go getUserInput(answerCh, scanner)

		select {
		case <-end:
			fmt.Println("\nTime's Up!")
			return correct
		case answer := <-answerCh:
			if answer == row[1] {
				correct++
			}
		}
	}

	return correct
}

func getUserInput(answerCh chan string, scanner *bufio.Scanner) {
	scanner.Scan()
	answer := strings.ToLower(strings.Trim(scanner.Text(), " "))

	answerCh <- answer
}
