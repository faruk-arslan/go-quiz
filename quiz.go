package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {

	file, err := os.Open("questions.csv") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(records)
	// fmt.Println(len(records))
	// fmt.Println(reflect.TypeOf(records))

	// rt := reflect.TypeOf(records)
	// switch rt.Kind() {
	// case reflect.Slice:
	// 	fmt.Println(records, "is a slice with element type", rt.Elem())
	// case reflect.Array:
	// 	fmt.Println(records, "is an array with element type", rt.Elem())
	// default:
	// 	fmt.Println(records, "is something else entirely")
	// }

	var point float32 = 0
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	done := make(chan bool)

	for i := 0; i < len(records); i++ {
		var tempAnswer int

		go func() {
			fmt.Printf("Question %d: %v=? \n", i+1, records[i][0])
			fmt.Scan(&tempAnswer)

			// convert the answer string (from slice) to int
			intAnswer, err := strconv.Atoi(records[i][1])
			if err != nil {
				log.Fatal("Can not convert the answer to int")
			}
			if tempAnswer == intAnswer {
				point += 8.333333333333333
			}
			// continue to next iteration
			done <- true
		}()

		select {
		case <-done:
			// if done true, continue to the next iteration
			fmt.Println("Next question!")
			continue
		case t := <-ticker.C:
			// ticker.C is triggered after timeout
			fmt.Println("Current time: ", t)
			fmt.Println("Time is out, next question: ")
		}
	}
	fmt.Println("Your score is: ", point)
}
