package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	//Opens CSV file with questions in it
	file, err := os.Open("problems.csv")

	//Error handling
	if err != nil {
		log.Fatal(err)
	}
	//Waits until function is complete to close connection to file
	defer file.Close()

	//Creates a reader to extract information from file
	csv := csv.NewReader(file)

	// Stores questions in a 2d slice
	questions, err := csv.ReadAll()

	//Error handling
	if err != nil {
		log.Fatal(err)
	}
	// Track user score
	score := 0

	//* Version 1
	// // Ask user the questions. If they are correct, their score goes up
	// for qNum, question := range questions {
	// 	//Ask question
	// 	fmt.Printf("Question %d: what is %s? ", qNum+1, question[0])
	// 	//Wait for user input
	// 	reader := bufio.NewReader(os.Stdin)
	// 	input, _ := reader.ReadString('\n')

	// 	//Compare user input to answer
	// 	if strings.TrimSpace(input) == question[1] {
	// 		fmt.Println("Correct!")
	// 		score++
	// 	} else {
	// 		fmt.Println("Incorrect!")
	// 		break
	// 	}
	// }

	// Create slice of problem structs
	questionList := parseProblems(questions)

	// Don't start timer until user is ready
	fmt.Printf("When ready, press enter ")
	fmt.Scanf("%s")

	for i, question := range questionList {
		fmt.Printf("Question #%d, what is %s? ", i+1, question.q)

		var answer string
		fmt.Scanf("%s", &answer)

		if answer == question.a {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Println("Incorrect!")
			break
		}
	}

	fmt.Printf("Your score is %d\n", score)
}

// Struct to hold problem question and answer
type problem struct {
	q string
	a string
}

func parseProblems(problems [][]string) []problem {
	problemSlice := make([]problem, len(problems))
	for i, line := range problems {
		problemSlice[i] = problem{q: line[0], a: line[1]}
	}
	return problemSlice
}
