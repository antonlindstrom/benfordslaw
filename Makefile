all: build

build:
	@mkdir -p bin/
	@go build -o bin/example example/example.go

get-deps:
	@go get -d -v ./...

test: build
	@go test ./...

coverage:
	@go get code.google.com/p/go.tools/cmd/cover
	@go test -cover ./...

benchmark:
	@go test -bench ./...

clean:
	@rm -rf bin/
