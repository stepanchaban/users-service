run:
	go run cmd/server/main.go

lint:
	golangci-lint run --color=always