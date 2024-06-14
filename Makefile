build:
	go build -o ./bin/grgrep ./cmd/grgrep

test: build
	go test ./...