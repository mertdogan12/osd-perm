package conf

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var Port int = 80

func Parse(args []string) {
	for i, arg := range args {
		switch arg {
		case "-p":
			port, err := strconv.Atoi(args[i+1])
			if err != nil {
				log.Fatal("Given port is not a number")
				return
			}
			Port = port
			return

		case "--help":
			fmt.Println(
				"Arguments:\n" +
					"     -p     Port\n" +
					"     --help Shows the help menu")
			os.Exit(0)

		default:
			log.Fatalf("Argument %s does not exists. --help to the all commands", arg)
			return
		}
	}
}
