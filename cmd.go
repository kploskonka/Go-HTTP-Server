package main

import (
    "fmt"
    "os"
    "strings"
)

const version = "1.0.0"

func printHelp() {
    fmt.Print("===== Available commands =====\n" +
    "help - Displays all commands\n" +
    "version - Displays the current version of the program\n" +
    "run --file <index.html> - Starts an HTTP web server serving the HTML file provided as an argument\n" +
    "==============================\n")
}

func printVersion() {
    fmt.Printf("Current version: %v\n", version)
}

func printUnknownCmd() {
    fmt.Println("Unknown arguments. Execute with 'help' to get available commands.")
}

func HandleTextCommand(command string) {
    switch command {
    case "help":
        printHelp()
    case "version":
        printVersion()
    default:
        printUnknownCmd()
    }
}

func exists(fileName string) bool {
    if _, err := os.Stat(fileName); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }

    return true
}

func IsRunCommandValid(args []string, fileName string) bool {
    if !strings.Contains(args[2], "-file") {
        printUnknownCmd()
        return false;
    }

    if !exists(fileName) {
        fmt.Printf("File '%v' does not exist.\n", fileName)
        return false
    }

    return true
}