package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	filename := flag.String("file", "quiz_questions.csv", "")
	timeInSeconds := flag.Int("time", 30, "")
	shuffle := flag.Bool("shuffle", false, "")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading records from file", err)
	}

	correctAnswers := 0
	totalQuestions := len(records)
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Printf("You have %d seconds for this test\n", *timeInSeconds)
	fmt.Println("Press Enter to START the QUIZ...")
	inputReader.ReadString('\n')

	timer := time.NewTimer(time.Duration(*timeInSeconds) * time.Second)
	done := make(chan bool)

	if *shuffle {
		rand.New(rand.NewSource(time.Now().UnixNano()))
		rand.Shuffle(len(records), func(i, j int) {
			records[i], records[j] = records[j], records[i]
		})
	}

	go func() {
		for i, row := range records {
			question := row[0]
			options := strings.Split(row[1], ",")
			answer := row[2]

			fmt.Println(fmt.Sprintf("Q.%d.", i+1), question)
			for i, option := range options {
				fmt.Printf("%v) %v\n", i+1, strings.TrimSpace(option))
			}
			fmt.Print("Your Answer: ")
			input, _ := inputReader.ReadString('\n')
			input = strings.TrimSpace(input)

			if strings.ToLower(answer) == strings.ToLower(input) {
				correctAnswers = correctAnswers + 1
			}
			fmt.Println()
		}
		done <- true
	}()

	select {
	case <-timer.C:
		fmt.Println("\n\nTime's up")
	case <-done:
		fmt.Println("\nOn time")
	}

	fmt.Printf("Your score is %d/%d", correctAnswers, totalQuestions)
}
