package main

import (
	"fmt"
	"os"
	"bufio"
	"encoding/csv"
	"flag"
	"time"
)

var limit int
var file string

func init() {
	flag.IntVar(&limit, "limit", 1, "limit time")
	flag.StringVar(&file, "file", "./problems.csv", "file path")
}

func main()  {
	flag.Parse()
	printStats(quiz())
	os.Exit(0)
}

func quiz() (score int , total int){
	r := csv.NewReader(getReader(file))
	record, err := r.Read()
	for record != nil {
		if err != nil {
			exitWithError(err.Error(), 1)
		}
		total ++
		fmt.Println(record[0])
		if query(record[1]) {
			score ++
		}
		record, err = r.Read()
	}
	return score, total
}

func query(solution string) (isCorrect bool) {
	channel := make(chan string)
	go func () {
		var input string
		fmt.Scanln(&input)
		channel <-input
	}()
    select{
	case res := <-channel:
		return res == solution
	case <-time.After(time.Duration(limit) * time.Second):
		return false
	}
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