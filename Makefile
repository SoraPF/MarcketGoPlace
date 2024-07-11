build:
	@go build -o bin/Marcketplace .

run: build
	./bin/Marcketplace