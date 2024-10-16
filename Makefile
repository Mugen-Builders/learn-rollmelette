-include .env.develop

START_LOG = @echo "================================================= START OF LOG ==================================================="
END_LOG = @echo "================================================== END OF LOG ===================================================="

.PHONY: tests
tests:
	@go test -p 1 ./... -coverprofile=./coverage.md -v

.PHONY: machine
machine:
	$(START_LOG)
	@docker build \
		-t machine:latest \
		-f ./build/Dockerfile .
	@cartesi build --from-image machine:latest
	$(END_LOG)