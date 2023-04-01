.PHONY: all
all: build

.PHONY: build
build: clean
	mkdir -p bin
	echo "build html..."
	cd html && npm i && npm run build
	echo "build server..."
	go build -o ./bin/server

.PHONY: server
server:
	go build -o ./bin/server

.PHONY: run
run:
	./bin/server

.PHONY: html
html:
	cd html && npm start

.PHONY: clean
clean:
	rm -rf bin

.PHONY: help
help:
	@echo "make all - build server && html"
	@echo "make run - serve at localhost:8080"
