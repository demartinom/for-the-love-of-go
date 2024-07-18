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

	//Placeholder
	fmt.Println(questions)
}
