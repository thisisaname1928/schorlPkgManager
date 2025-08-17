GO_SRC=$(shell find . -type f -name "*.go")
OUTPUT=./SchorlPackageManager

$(OUTPUT): $(GO_SRC)
	@echo "Building..."
	@go build

run: $(OUTPUT)
	@./$< selfInstall TestingFolder

