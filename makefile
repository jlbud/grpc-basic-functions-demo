build:
	protoc --go_out=plugins=grpc:./ proto/depend.proto
	protoc --go_out=plugins=grpc:./ proto/helloworld.proto

ignore-import-build:
	protoc --go_out=plugins=grpc:./ --go_opt=Mproto/depend.proto=proto proto/depend.proto
	protoc --go_out=plugins=grpc:./ --go_opt=Mproto/helloworld.proto=proto proto/helloworld.proto

build-all:
	protoc --go_out=plugins=grpc:./ proto/helloworld.proto
	protoc-go-inject-tag -input=proto/helloworld.pb.go

start-server:
	go test -v -run TestClientStreamServer

start-client:
	# client stream
	go test -v -run TestClientStreamCli