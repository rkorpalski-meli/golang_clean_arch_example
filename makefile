install-mocks:
	@go get github.com/vektra/mockery/.../

build-mocks:
	@~/go/bin/mockery -dir src/api -output src/api/mocks -recursive -all

test:
	@go test ./...

linter-install:
	@go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.27.0

lint:
	@~/go/bin/golangci-lint run

fmt:
	@go fmt ./...