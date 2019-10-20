package main

import (
	"bufio"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/net/html"
	"os"
)

func main() {
	nodes, err := html.Parse(getReader("./ex2.html"))
	if err != nil {
		panic(err)
	}
	spew.Dump(nodes)

}

func getReader(path string) *bufio.Reader {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return bufio.NewReader(file)
}
