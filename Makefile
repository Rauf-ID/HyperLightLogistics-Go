generate:
	mkdir pkg/service
	protoc \
		--proto_path=api/proto "api/proto/*.proto" \
		--go_out=pkg/service --go_opt=paths=source_relative \
		--go-grpc_out=pkg/service --go-grpc_opt=paths=source_relative