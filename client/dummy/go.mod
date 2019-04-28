module main

go 1.12

require (
	github.com/LMJW/werewolves/server/proto v0.0.0
	github.com/golang/protobuf v1.3.1 // indirect
	google.golang.org/grpc v1.20.1
)

replace github.com/LMJW/werewolves/server/proto => ../../server/proto
