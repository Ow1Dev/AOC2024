# List available commands
default:
    @just --list

# Run solution for a giving day 
run day:
    @go run {{day}}/main.go

# Run tests for a giving day 
test day:
    @go test ./{{day}}/...


# Run test all solutions
test-all:
    @go test ./...


# format the code
fmt:
    @go fmt ./...
