package main

import (
	"flag"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
	"strings"
)

func main() {
	// Use flag to show a hint to the user when used the -h flag
	fileName := flag.String("csv","problems.csv","csv file to read the questions and answers")
	timeLimit := flag.Int("time", 30, "time limit in seconds")
	flag.Parse()

	file, err := os.Open(*fileName) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// Make a slice which contains the problems in a 'problem' struct type.
	problems := makeStructSlice(records)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	done := make(chan bool)
	correctAnswers := 0

	for i, p := range problems {
		var tempAnswer string

		go func() {
			fmt.Printf("Question %d: %v=? \n", i+1, p.q)
			fmt.Scan(&tempAnswer)
			if(tempAnswer == p.a){
				correctAnswers++
			}
			done <- true
		}()

		select {
		case <-done:
			// if done true, continue to the next iteration
			fmt.Println("Next question!")
			timer = time.NewTimer(time.Duration(*timeLimit) * time.Second)
			continue
		case t := <-timer.C:
			fmt.Println("Current time: ", t)
			fmt.Println("Time is out, next question: ")
			timer = time.NewTimer(time.Duration(*timeLimit) * time.Second)
 		}
	}
	fmt.Printf("You scored %d out of %d", correctAnswers, len(records))
}

// Define a struct for the problems.
type problem struct{
	q string
	a string
}

func makeStructSlice(records [][]string) []problem{
	// Make a slice which contains the problems in a 'problem' struct type.
	problems := make([]problem, len(records))
	for i,record := range records{
		// If we know the length, it's more efficient than append.
		problems[i] = problem{q:record[0], a:strings.Trim(record[1]," ")}
	}
	return problems
}
