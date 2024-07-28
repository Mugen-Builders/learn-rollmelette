.PHONY: tests
tests:
	@go test -p 1 ./... -coverprofile=./coverage.md -v