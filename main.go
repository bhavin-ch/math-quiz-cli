package main

import (
	"strings"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	question string
	answer string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A csv file with questions and answers")
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
	problems := parseLines(lines)
	counter := 0
	for i, prob := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, prob.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == prob.answer {
			counter++
		}
	}
	fmt.Printf("You scored %d out of %d problems correctly\n", counter, len(problems))
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