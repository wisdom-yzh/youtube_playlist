.PHONY: all
all: build start

.PHONY: build
build: clean
	mkdir -p bin
	echo "build server..."
	go build -o ./bin/server
	echo "build html..."
	cd html && npm i && npm run build && mv build ../bin/html

.PHONY: run
run:
	./bin/server

.PHONY: html
html:
	cd html && npm start

.PHONY: clean
clean:
	rm -rf bin

.PHONY: test
test:
	@cd server
	@go test

.PHONY: help
help:
	@echo "make all - build server"
