PRG := unsavory

run: build
	@./bin/${PRG} --dry-run -token ${PINBOARD_TOKEN}

build:
	@go build -o bin/${PRG} ./cmd/${PRG}

lint:
	@go vet ./...
	@staticcheck ./...
	@shadow ./...

clean:
	@go clean
	@rm bin/${PRG}

.PHONY: build run lint clean
