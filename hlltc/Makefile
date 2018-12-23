default: clean generate

generate:
	protoc *.proto --go_out=plugins=grpc:.

clean:
	rm -f *.pb.go
