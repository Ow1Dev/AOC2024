# List available commands
default:
    @just --list

run day part:
    @go run {{day}}/{{part}}/main.go

fmt:
    @go fmt ./...
