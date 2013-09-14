all:
	@mkdir -p bin/
	@go build -o bin/benford benford.go

get-deps:
	@go get -d -v ./...

test:
	@go test ./...

clean:
	@rm -rf bin/
