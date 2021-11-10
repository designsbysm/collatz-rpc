.PHONY: all

default: test lint


# coverage:
# 	go test ./... -coverprofile=coverage.out
# 	go tool cover -html="coverage.out"

# # docs:
# 	godoc2md github.com/designsbysm/server-go | sed -e s#src/target/##g > DOCUMENTATION.md

lint:
	golangci-lint run

proto: collatz.proto
	protoc --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import collatz.proto
	
test:
	go test ./... -cover -race
