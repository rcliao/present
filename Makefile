all: clean build run

clean:
	rm -rf bin/main

build: clean
	go build -o bin/main cmd/main.go

run: clean
	./bin/main
