# Go-HTTP-Server
Simple command-line interface program which runs an HTTP web server with a file provided as a CLI argument.

## Installation guide
```
git clone https://github.com/kploskonka/Go-HTTP-Server.git
cd GO-HTTP-Server
go test . (optional, to check if tests are passing)
go build
.\Go-HTTP-Server.exe <command> <arguments>
```
Alternatively, go to [Releases](https://github.com/kploskonka/Go-HTTP-Server/releases/tag/1.0.0), download `Go-HTTP-Server.exe` and run the last command.


## Available commands
```
help - Displays all commands
version - Displays the current version of the program
run --file <index.html> - Starts an HTTP web server serving the HTML file provided as an argument
```
___
Katarzyna Płoskonka in 2021
