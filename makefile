proto:
	protoc -I server/proto server/proto/werewolf.proto --go_out=plugins=grpc:server/proto

clean:
	rm -rf server/proto/*.pb.go