all:
	@mkdir -p bin/
	@go build -o bin/example example/example.go

get-deps:
	@go get -d -v ./...

test:
	@go test -v ./...

clean:
	@rm -rf bin/
