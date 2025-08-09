LOCAL_DIR=$(CURDIR)
LOCAL_BIN=$(LOCAL_DIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint:
	$(LOCAL_BIN)/golangci-lint --version && $(LOCAL_BIN)/golangci-lint run -v

run:
	go run cmd/app/main.go