package main

import (
	"fmt"
	"github.com/YoSmudge/server-names/namer"
	"io/ioutil"
	"net/http"
	"os"
)

const instanceIdUrl string = "http://169.254.169.254/latest/meta-data/instance-id"

func main() {
	var serverId string

	if len(os.Args) >= 2 {
		serverId = os.Args[1]
	}

	if serverId == "" {
		var ra []byte
		r, err := http.Get(instanceIdUrl)

		if err == nil {
			ra, err = ioutil.ReadAll(r.Body)
		}

		if err != nil {
			fmt.Println("Instance ID not supplied and could not talk to the instance metadata service:", err)
			os.Exit(1)
		}
		serverId = string(ra)
	}

	name, err := namer.Name(serverId)
	if err != nil {
		fmt.Println("Error generating name:", err)
		os.Exit(1)
	}
	fmt.Println(name)
}
