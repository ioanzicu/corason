.PHONY: test
test:
	go clean -testcache
	go test -race -v -cover ./...

.PHONY: run
run:
	go run ./cmd/main.go