# List available commands
default:
    @just --list

run day:
    @go run {{day}}/main.go

test day:
    @go test ./{{day}}/...

test-all:
    @go test ./...

fmt:
    @go fmt ./...
