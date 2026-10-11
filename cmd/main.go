package main

import (
	"flag"
	"fmt"
	command "openeff/intera/cmd/internal"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: intera help")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		createCmd := flag.NewFlagSet("create", flag.ExitOnError)

		output := createCmd.String("output", "nil", "Output path")

		createCmd.Parse(os.Args[2:])

		fmt.Println("Output:", *output)

	case "help":
		fmt.Println(command.Help())

	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
}
