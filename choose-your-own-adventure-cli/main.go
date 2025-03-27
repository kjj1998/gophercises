package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	models "github.com/kjj1998/gophercises/choose-your-own-adventure-cli/models"
)

var mappings models.Story
var currentArc string = "intro"

func main() {
	file, err := os.Open("data/gopher.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	_ = decoder.Decode(&mappings)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if currentArc == "home" {
			break
		}
		title := mappings[currentArc].Title
		story := mappings[currentArc].Story
		options := mappings[currentArc].Options

		fmt.Printf("Title: %s\n", title)

		for _, v := range story {
			fmt.Printf("%s\n", v)
		}
		fmt.Println()

		for i, v := range options {
			fmt.Printf("%s (Press %d)\n", v.Text, i+1)
		}

		input, _ := strconv.Atoi(getUserInput(scanner))

		currentArc = options[input-1].Arc
		fmt.Println()
	}

	endTitle := mappings[currentArc].Title
	ending := mappings[currentArc].Story

	fmt.Printf("Title: %s\n", endTitle)

	for _, v := range ending {
		fmt.Printf("%s\n", v)
	}
	fmt.Println()
}

func getUserInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	answer := strings.ToLower(strings.Trim(scanner.Text(), " "))

	return answer
}
