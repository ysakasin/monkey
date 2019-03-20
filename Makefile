.PHONY: monkey
monkey:
	go build monkey

.PHONY: run
run:
	go run main.go

.PHONY: test
test:
	go test ./...
