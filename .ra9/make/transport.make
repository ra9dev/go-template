PROTOC := $(shell command -v protoc 2> /dev/null)
PROTO_PATH ?= ./pb/*.proto
grpc:
ifndef PROTOC
	$(error "protoc is not installed. Visit https://grpc.io/docs/protoc-installation")
endif
	$(call describe_job,"Generating grpc code")
	protoc \
    	--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go --go_out=. --go_opt=paths=source_relative --go_opt=M$(PROTO_PATH)=./pb \
    	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc --go-grpc_out=. --go-grpc_opt=paths=source_relative --go-grpc_opt=M$(PROTO_PATH)=./pb \
    	-I /usr/local/include:$(THIRD_PARTY_PROTO_PATH):. \
    	$(PROTO_PATH)