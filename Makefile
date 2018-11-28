all: build

generate:
	go generate ./...
	protoc -I services/grpc/ services/grpc/service.proto --go_out=plugins=grpc:services/grpc

build:
	go build -v -o service .

run_unit_tests:
	go test -run 'Unit' -v -race ./...
	go test -run 'Unit' -v -covermode=count -coverprofile=./coverage.out ./...

clean:
	rm -rf ./service
	rm -rf ./coverage.out
