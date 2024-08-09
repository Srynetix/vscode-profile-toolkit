_all:
    @just --list

# Format
fmt:
    go fmt ./...

# Lint
lint:
    go vet ./...

# Test
test:
    go test ./...

# Run
run *ARGS:
    go run main.go {{ ARGS }}

# Build
build:
    go build -o vs-prof-tk