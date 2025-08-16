GO_SRC=$(shell find . -type f -name "*.go")
OUTPUT=./TestingFolder/app/org.schorl.rlg/bin/rlg

$(OUTPUT): ./SchorlPackageManager
	@cp $< $@

run: $(OUTPUT)
	@$< init

./SchorlPackageManager: $(GO_SRC)
	@echo "Building..."
	@go build