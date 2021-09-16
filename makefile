build:
	protoc --go_out=plugins=grpc:./ proto/depend.proto
	protoc --go_out=plugins=grpc:./ proto/helloworld.proto

build-all:
	protoc --go_out=plugins=grpc:./ proto/helloworld.proto
	protoc-go-inject-tag -input=proto/helloworld.pb.go