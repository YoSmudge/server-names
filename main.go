package main

import (
	"fmt"
	"github.com/YoSmudge/server-names/namer"
	"os"
)

func main() {
	serverId := os.Args[1]

	name, err := namer.Name(serverId)
	if err != nil {
		fmt.Println("Error generating name:", err)
		os.Exit(1)
	}
	fmt.Println(name)
}
