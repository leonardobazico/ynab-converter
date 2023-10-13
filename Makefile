install-ci-dependencies:
	go version
	go install gotest.tools/gotestsum
	go install github.com/securego/gosec/cmd/gosec

setup-dev: install-ci-dependencies
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54

lint:
	golangci-lint run

security-check:
	gosec ./...

test: lint security-check
	gotestsum --format=testname ./...

test-watch:
	gotestsum --format=testname --watch ./...

test-coverage:
	mkdir -p output
	gotestsum --format=testname -- -coverprofile=output/coverage.out ./...
