GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

.PHONY: init
# init env
init:
	go install github.com/rakyll/statik@v0.1.7

.PHONY: clean
# clean
clean:
	rm -rf ./target

.PHONY: generate
# generate
generate:
	statik -src=./res/raw -dest=./res -f -include=* -ns=res
	go fmt ./...

.PHONY: build
# build
build:
	go build -trimpath -ldflags "-X main.Version=$(VERSION)" -o ./target/loggen ./

.PHONY: all
# all
all:
	make clean;
	make generate;
	make build;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
