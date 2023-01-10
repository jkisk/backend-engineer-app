# Backend-Engineer-App

## Requirements
1. `Go 1.19`

## Installation & operation
1. Clone repo onto your machine with GitHub CLI `gh repo clone jkisk/backend-engineer-app`
1. Open a terminal and navigate to the directory containing backend-engineer-app 
1. Run `go mod tidy`
1. Optionally configure port by running `export BEA_PORT="xxxx"`
1. Run service locally by running `go run main.go`
1. Test service by running `curl localhost:$PORT/employees`