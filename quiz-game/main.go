package main

import (
	"bufio"
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

	correct, wrong := quiz(&data)

	fmt.Printf("Number of questions answered correctly: %d\n", correct)
	fmt.Printf("Number of questions answered wrongly: %d\n", wrong)
}

func quiz(questionsAndAnswers *[][]string) (correct, wrong int) {
	correct = 0
	wrong = 0

	scanner := bufio.NewScanner(os.Stdin)

	for index, row := range *questionsAndAnswers {
		fmt.Printf("Problem #%d: %s = ", index+1, row[0])

		scanner.Scan()
		err := scanner.Err()

		if err != nil {
			panic(err)
		}

		answer := scanner.Text()

		if answer == row[1] {
			correct++
		} else {
			wrong++
		}
	}

	return correct, wrong
}
