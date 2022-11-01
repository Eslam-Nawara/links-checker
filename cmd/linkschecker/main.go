package main

import (
	"fmt"
    "github.com/Eslam-Nawara/linkschecker"
)

func main() {
	err := linkschecker.CheckLinksInFile()

	if err != nil {
		fmt.Println(err)
	}
}
