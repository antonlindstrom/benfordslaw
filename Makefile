all:
	@mkdir -p bin/
	@go build -o bin/benford benford.go

test:
	@go test ./...

clean:
	@rm -rf bin/
