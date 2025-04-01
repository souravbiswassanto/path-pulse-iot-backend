SRC_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
PROTO_DIR:=$(SRC_DIR)/proto
OUT_DIR:=$(SRC_DIR)/protogen/golang

protoc: protoc-pb gwprotos

protoc-pb:
	protoc --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
	--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
	--proto_path=$(PROTO_DIR) 			\
	$(PROTO_DIR)/**/**/*.proto

prev-protoc:
	cd proto && protoc --go_out=../protogen/golang --go_opt=paths=source_relative \
	--go-grpc_out=../protogen/golang --go-grpc_opt=paths=source_relative \
	./**/**/*.proto

run:
	rm -f path-pulse-iot-backend && go build -o path-pulse-iot-backend .

gwprotos:
	echo "Generating gRPC Gateway bindings and OpenAPI spec"
	protoc -I . --grpc-gateway_out $(OUT_DIR) 								\
    --grpc-gateway_opt logtostderr=true 									\
    --grpc-gateway_opt paths=source_relative 									\
    --grpc-gateway_opt generate_unbound_methods=true 								\
    --proto_path=$(PROTO_DIR) 									\
      $(PROTO_DIR)/**/**/*.proto