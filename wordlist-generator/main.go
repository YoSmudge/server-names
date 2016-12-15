package main

import (
	"fmt"
	"github.com/YoSmudge/server-names/wordlist"
	"github.com/voxelbrain/goptions"
	"os"
)

type options struct {
	Help   goptions.Help `goptions:"-h, --help, description='Show help'"`
	Source string        `goptions:"-s, --source, description='Source directory for WordNet data'"`
	Dest   string        `goptions:"--dest, description='Destination for word files'"`
}

func main() {
	parsedOptions := options{}

	goptions.ParseAndFail(&parsedOptions)

	fmt.Println("Generating wordlist")
	err := wordlist.Generate(parsedOptions.Source, parsedOptions.Dest)
	if err != nil {
		fmt.Println("Error generating wordlist:", err)
		os.Exit(1)
	}
}
