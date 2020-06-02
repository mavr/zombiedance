build:
	CGO_ENABLED=0 go build -o bin/zombiedance cmd/zombiedance/*.go

run: build
	./bin/zombiedance
