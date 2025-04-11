GOTEST=gotestsum --format=testname

install-ci-dependencies:
	go version
	go install github.com/securego/gosec/v2/cmd/gosec@v2.20.0
	go install github.com/go-critic/go-critic/cmd/gocritic@v0.11.4


setup-dev: install-ci-dependencies
	go install gotest.tools/gotestsum@v1.12.0
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.0.2

lint:
	golangci-lint run ./...

lint-fix:
	golangci-lint run --fix ./...

security-check:
	gosec ./...

critic:
	gocritic check -enableAll ./...

test: lint security-check critic
	$(GOTEST) -- -count=1 ./...

test-watch:
	$(GOTEST) --watch ./...

test-integration:
	$(GOTEST) -- -count=1 ./cmd/...

test-ci:
	rm -rf coverage/*
	mkdir -p coverage/integration
	mkdir -p coverage/unit
	SKIP_INTEGRATION=true \
		go test -cover -count=1  ./... -args -test.gocoverdir="$(PWD)/coverage/unit"
	GOCOVERDIR=coverage/integration \
		go test -count=1 ./cmd/...
	make coverage-files

coverage-files:
	go tool covdata textfmt -i=./coverage/unit,./coverage/integration -o coverage/profile.out
	sed -i'.bak' "s#$(PWD)#ynabconverter#g" coverage/profile.out
	go tool cover -func coverage/profile.out
	go tool cover -html=coverage/profile.out -o=coverage/profile.html

add-pre-commit-hook:
	rm -f .git/hooks/pre-commit
	ln -s -f ../../scripts/pre-commit.sh .git/hooks/pre-commit

prettier:
	npx prettier '**/*.{yml,md}' --write

build-cli:
	rm -rf bin/*
	go build -o bin/ ./cmd/...

test-pipeline:
	act -s GITHUB_TOKEN="$(gh auth token)"
