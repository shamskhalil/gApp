create:
	protoc --proto_path=proto/ proto/*.proto --go_out=gen
	protoc --proto_path=proto/ proto/*.proto --go-grpc_out=gen
clean:
	rm -r gen/contactpb/*.go