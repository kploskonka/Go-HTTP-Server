# Go-HTTP-Server
Simple command-line interface program which runs an HTTP web server with a file provided as a CLI argument.

Usage:
```
git clone https://github.com/kploskonka/Go-HTTP-Server.git
go build
.\Go-HTTP-Server.exe <command> <arguments>
```

Available commands:
```
help - Displays all commands
version - Displays the current version of the program
run --file <index.html> - Starts an HTTP web server serving the HTML file provided as an argument
```