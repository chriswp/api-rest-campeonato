build:
	@go build -o build/api-rest-campeonato cmd/main.go

test:
	@go test -v ./...

run: build
	@./build/api-rest-campeonato