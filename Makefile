build:
	mkdir -p bin
	GOOS=linux go build -o bin/aub

run:
	./bin/aub