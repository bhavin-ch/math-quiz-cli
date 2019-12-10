package main

import (
	"strings"
	"encoding/csv"
	"flag"
	"time"
	"fmt"
	"os"
)

type problem struct {
	question string
	answer string
}

func main() {
	defaultCsvPath := "problems.csv"
	defaultTimeLimit := 30
	lines, limit := parseFlags(&defaultCsvPath, &defaultTimeLimit)
	timer := time.NewTimer(time.Duration(limit) * time.Second)
	problems := parseLines(lines)
	askQuestions(&problems, timer, &limit)
}

func parseFlags(defaultPath *string, defaultTime *int) ([][]string, int) {
	csvFilename := flag.String("csv", *defaultPath, "A csv file with questions and answers")
	limit := flag.Int("limit", *defaultTime, "Time limit for the quiz")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed the parse the csv file")
	}
	return lines, *limit
}

func askQuestions(problems *[]problem, timer *time.Timer, limit *int) {
	counter := 0
	problemLoop: for i, prob := range *problems {
		fmt.Printf("Problem #%d: %s = ", i+1, prob.question)
		ansCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			ansCh <- answer
		}()
		select {
			case <-(*timer).C:
				fmt.Printf("(Your %d secs are up!)\n", *limit)
				break problemLoop
			case answer := <-ansCh:
				if answer == prob.answer {
					counter++
				}
		}
	}
	fmt.Printf("You answered %d out of %d problems correctly\n", counter, len(*problems))
}

func parseLines(lines [][]string) []problem {
	ans := make([]problem, len(lines))
	for i, line := range lines {
		ans[i] = problem{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}
	return ans
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}