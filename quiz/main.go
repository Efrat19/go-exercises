package main

import (
	"fmt"
	"os"
	"bufio"
	"encoding/csv"
	"flag"
)

var limit int
var file string

func init() {
	flag.IntVar(&limit, "limit", 5, "limit time")
	flag.StringVar(&file, "file", "./problems.csv", "file path")
}

func main()  {
	flag.Parse()
	quiz()
}

func quiz(){
	var score int = 0
	var total int = 0

	r := csv.NewReader(getReader(file))
	record, err := r.Read()
	for record != nil {
		if err != nil {
			exitWithError(err.Error(), 1)
		}
		total ++
		if query(record[0],record[1]) {
			score ++
		}
		record, err = r.Read()
	}
	printStats(score, total)
	os.Exit(0)
}

func query(exp string, solution string) (isCorrect bool) {
    var input string
	fmt.Println(exp)
    fmt.Scanln(&input)
	return input == solution
}

func getReader(path string) *bufio.Reader {
	file, err := os.Open(path)
	if err != nil {
		exitWithError(err.Error(), 1)
	}
	return bufio.NewReader(file)
}

func printStats(score int, total int)  {
	err := fmt.Sprintf("you scored %d out of %d", score, total)
	if err != "" {
		exitWithError(err, 1)
	}
}

func exitWithError(msg string, code int) {
	fmt.Println(msg)
	os.Exit(code)
}