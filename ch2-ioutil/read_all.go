package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ReadFile和WriteFile

// ReadAll

func TestReadAll() {
	f, err := os.Open("./a.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	p, err := ioutil.ReadAll(f)
	fmt.Printf("%s", p)
}

func main() {
	TestReadAll()
}
