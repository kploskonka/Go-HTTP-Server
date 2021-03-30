package main

import "fmt"

const version = "1.0.0"

func PrintHelp() {
	fmt.Print("===== Available commands =====\n" +
	"help - Displays all commands\n" +
	"version - Displays the current version of the program\n" +
	"run --file <index.html> - Starts an HTTP web server serving the HTML file provided as an argument\n" +
	"==============================\n")
}

func PrintVersion() {
	fmt.Printf("Current version: %v\n", version)
}

func PrintUnknownCmd() {
	fmt.Println("Unknown arguments. Execute with 'help' to get available commands.")
}