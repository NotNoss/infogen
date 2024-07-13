package main

import (
	"fmt"
	"os"

	"github.com/NotNoss/infogen/app"
	"github.com/NotNoss/infogen/util"
)

const help = "Email is provided by default if there is no argument provided.\n\n-n, --name  Pull name from API\n-a, --address  Pull adress from API\n-c, --config  Generate a config file\n-h, --help  This help screen you are seeing"

func main() {
	args := os.Args

	if len(args) < 2 {
		app.Email()
	}

	switch args[1] {
	case "-n", "--name":
		app.Name()
	case "-m", "--mail":
		app.Email()
	case "-a", "--address":
		app.Address()
	case "-c", "--config":
		util.CheckConfig()
	case "-h", "--help":
		fmt.Println(help)
	default:
		fmt.Println(help)
	}
}
