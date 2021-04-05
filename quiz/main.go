package main

// Importing go packages
import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Creating a strcut
type problem struct {
	qus string
	ans string
}

//  Main function
func main() {
	// csv file name
	csv_file := flag.String("csv", "questions_math.csv", "something")
	// timeout limit
	timelimit := 5
	flag.Parse()
	// Opening and reading the csv file
	file, _ := os.Open(*csv_file)
	r := csv.NewReader(file)
	lines, _ := r.ReadAll()

	//  using a functions to add csv table values to slice
	ques := parseLines(lines)
	// Counting number of correct answers by the user
	correct := 0
	// timer module to time out
	timer := time.NewTimer(time.Duration(timelimit) * time.Second)
problemloop:
	for i, p := range ques {
		fmt.Printf("Question #%d: what is %s, sir?\n", i+1, p.qus)
		// creating a answere channel
		answerCh := make(chan string)
		//  anonimus function
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		// case if timer runs out break the loop and return result
		case <-timer.C:
			fmt.Println()
			break problemloop
		// case if answere is provided keep running the loop
		case answer := <-answerCh:
			if answer == p.ans {
				correct++
			} else if answer == "exit" {
				fmt.Println("exited")
				return
			}
		}
	}
	//  final scorecard
	fmt.Printf("You scored %d out of %d.\n", correct, len(ques))
}

func parseLines(lines [][]string) []problem {
	//  creating a slice with struct type and length of total number of lines in csv
	ret := make([]problem, len(lines))
	//  adding values to slice
	for i, line := range lines {
		ret[i] = problem{
			qus: line[0],
			ans: strings.TrimSpace(line[1]),
		}
	}
	return ret
}
