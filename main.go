package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
)

type FileHandler struct {
    fileName string
}

func (fh *FileHandler) renderFile(res http.ResponseWriter, req *http.Request) {
    body, _ := ioutil.ReadFile(fh.fileName)

    fmt.Fprint(res, string(body))
}

func runServer(fileName string) {
    fmt.Printf("HTTP Server with %v is running on http://localhost:8090/. Press Ctrl+C to stop.\n", fileName)

    fileHandler := &FileHandler{fileName: fileName}

    http.HandleFunc("/", fileHandler.renderFile)
    http.ListenAndServe(":8090", nil)
}

func main() {
    runCommand := flag.NewFlagSet("run", flag.ExitOnError)
    fileName := runCommand.String("file", "", "file")

    if len(os.Args) < 2 {
        fmt.Println("Expected arguments. Execute with 'help' to get available commands.")
        return
    }

    if (strings.Contains(os.Args[1], "run")) {
        runCommand.Parse(os.Args[2:])

        if (IsRunCommandValid(os.Args, *fileName)) {
            runServer(*fileName)
        }
    } else {
        HandleTextCommand(os.Args[1])
    }
}