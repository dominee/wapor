.DEFAULT_GOAL=build

BINARY_NAME=wapor


clean:
	@echo "$$(tput bold)Cleaning$$(tput sgr0)"
	go clean
.PHONY:clean

test:
	@echo "$$(tput bold)Testing$$(tput sgr0)"
	go test ./... -v
.PHONY:test

test-coverage:
	@echo "$$(tput bold)Coverage test$$(tput sgr0)"
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out
.PHONY:test_coverage

dep:
	@echo "$$(tput bold)Downloading dependencies$$(tput sgr0)"
	go mod download
.PHONY:dep

vet: fmt
	@echo "$$(tput bold)Vetting$$(tput sgr0)"
	go vet
.PHONY:vet

fmt:
	@echo "$$(tput bold)Formating$$(tput sgr0)"
	go fmt ./...
.PHONY:fmt

lint: fmt
	@echo "$$(tput bold)Linting$$(tput sgr0)"
	golint ./...
.PHONY:lint

build: vet
	@echo "$$(tput bold)Building$$(tput sgr0)"
	go build -o ${BINARY_NAME} wapor.go
.PHONY:build

run: build
	@echo "$$(tput bold)Running...$$(tput sgr0)"
	cat testdata/urls.txt | ./${BINARY_NAME}
.PHONY:run