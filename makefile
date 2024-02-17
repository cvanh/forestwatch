BUILD_FLAGS=

build:
	go build $(BUILD_FLAGS) -o ./bin/main main.go    

run: build
	./bin/main

.PHONY: run