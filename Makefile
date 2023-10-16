GOTEST=gotestsum --format=testname

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
	$(GOTEST) ./...

test-watch:
	$(GOTEST) --watch ./...

test-coverage:
	mkdir -p output
	$(GOTEST) -- -coverprofile=output/coverage.out ./...
	go tool cover -html=output/coverage.out -o=output/coverage.html

add-pre-commit-hook:
	rm -f .git/hooks/pre-commit
	ln -s -f ../../scripts/pre-commit.sh .git/hooks/pre-commit

prettier:
	npx prettier '**/*.{yml,md}' --write
