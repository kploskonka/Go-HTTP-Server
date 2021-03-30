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

func exists(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }

    return true
}

func main() {
	runCommand := flag.NewFlagSet("run", flag.ExitOnError)
	fileName := runCommand.String("file", "", "file")

	if len(os.Args) < 2 {
		PrintUnknownCmd()
		return
	}
	
	switch os.Args[1] {
	case "help":
		PrintHelp()
	case "version":
		PrintVersion();
	case "run":
		runCommand.Parse(os.Args[2:])

		if (!exists(*fileName)) {
			fmt.Printf("File '%v' does not exist.\n", *fileName)
			return;
		}

		if strings.Contains(os.Args[2], "-file") {
			runServer(*fileName)
		} else {
			PrintUnknownCmd()
		}
	default:
		PrintUnknownCmd()
	}
}