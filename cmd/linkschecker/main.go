package main

import (
	"flag"
	"fmt"

	"github.com/Eslam-Nawara/linkschecker"
)

func main() {
	configFile := flag.String("config", "config.toml", "")
	flag.Parse()

	err := linkschecker.CheckLinksInFile(*configFile)

	if err != nil {
		fmt.Println(err)
	}
}
