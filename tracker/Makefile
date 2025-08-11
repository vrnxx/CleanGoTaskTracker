LOCAL_DIR=$(CURDIR)
LOCAL_BIN=$(LOCAL_DIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint:
	$(LOCAL_BIN)/golangci-lint --version && $(LOCAL_BIN)/golangci-lint run -v

build:
	go build -v cmd/app/main.go

run:
	go run cmd/app/main.go

code-lines:
	find . -name "*.go" -exec wc -l {} +
