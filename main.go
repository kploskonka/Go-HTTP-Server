package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const version = "0.0.1"

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
		runServer(*fileName)
	}
}