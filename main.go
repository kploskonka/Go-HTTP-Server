package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "0.0.1"

func main() {
	runCommand := flag.NewFlagSet("run", flag.ExitOnError)
	fileName := runCommand.String("file", "", "file")

	if len(os.Args) < 2 {
		fmt.Println("No arguments.")
		return
	}
	
	switch os.Args[1] {
	case "help":
		fmt.Println("Help command placeholder")
	case "version":
		fmt.Printf("Current version: %v\n", version)
	case "run":
		runCommand.Parse(os.Args[2:])
		fmt.Printf("Ran command 'run --file %v'\n", *fileName)
	}
}