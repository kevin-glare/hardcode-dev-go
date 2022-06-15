.PHONY:

PWD=$(shell pwd)

protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/messenger_proto/messenger.proto
