BUILD_FLAGS=

build:
	go build $(BUILD_FLAGS) -o ./bin/main main.go    

run: build
	./bin/main

lint:
	golangci-lint run

fmt:
	go fmt 

# run debug program
run_dbg: build
	dlv debug github.com/cvanh/forestwatch --headless --listen :2345 --log -- "$@"

package:
	fyne package -release -src ./

.PHONY: run, lint,run_dbg