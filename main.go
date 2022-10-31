package main

import (
	"fmt"
	linkschecker "linkschecker/links-checker"
)

func main() {
	err := linkschecker.CheckLinksInFile()
	if err != nil {
		fmt.Println(err)
	}
}
