package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	questionsAndAnswers := make(map[string]string)

	for _, row := range data {
		questionsAndAnswers[row[0]] = row[1]
	}

	quiz(&questionsAndAnswers)
}

func quiz(questionsAndAnswers *map[string]string) (correct, wrong int) {
	correct = 0
	wrong = 0

	for key, value := range *questionsAndAnswers {
		fmt.Println(key, " ", value)
	}

	return correct, wrong
}
