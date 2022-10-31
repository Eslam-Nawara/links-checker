package main

import (
	"fmt"
	"linkschecker"
)

func main() {
	err := linkschecker.CheckLinksInFile()

	if err != nil {
		fmt.Println(err)
	}
}
