BIN := "./bin/resizer"
DOCKER_IMG := "resizer:lastest"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/resizer

version:
	go $(BIN) -v

run: build
	$(BIN) -c ./configs/config.toml

build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .

run-img: build-img
	docker run --publish 2891:2891 $(DOCKER_IMG)

run-compose:
	docker-compose -f deployments/docker-compose.yml up -d resizer

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.50.1

lint: install-lint-deps
	golangci-lint run --skip-dirs docs ./...

test:
	go test -v -count=1 -race -timeout=1m ./tests/...

.PHONY: build version run build-img run-img run-compose lint test
