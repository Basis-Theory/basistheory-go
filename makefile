MAKEFLAGS += --silent

verify:
	go test ./...
