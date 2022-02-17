package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
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

	fmt.Println(records)
	fmt.Println(len(records))
	fmt.Println(reflect.TypeOf(records))

	rt := reflect.TypeOf(records)
	switch rt.Kind() {
	case reflect.Slice:
		fmt.Println(records, "is a slice with element type", rt.Elem())
	case reflect.Array:
		fmt.Println(records, "is an array with element type", rt.Elem())
	default:
		fmt.Println(records, "is something else entirely")
	}

	var point float32 = 0
	for i := 0; i < len(records); i++ {
		var tempAnswer int
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
	}
	fmt.Println("Your score is: ", point)

	// for i := 0; i < len(records); i++ {
	// 	fmt.Println(answers[i])
	// }

	// var jsonBlob = []byte(`[
	// 	{"Question": "5*2", "Answer": "10"}
	// ]`)

	// type Question struct {
	// 	Question string
	// 	Answer   string
	// }

	// var questions []Question
	// err := json.Unmarshal(data, &questions)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%+v", questions)
}
