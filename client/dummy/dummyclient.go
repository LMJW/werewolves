package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/LMJW/werewolves/server/proto"
	"google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("grpc dial error : %s \n", err)
	}

	dummyclient := pb.NewWerewolfClient(con)

	r, err := dummyclient.CreateID(context.TODO(), &pb.Signature{
		Username:    "player1",
		Displayname: "Groot",
		Randomcode:  123,
	})
	if err != nil {
		fmt.Printf("createID service error : %s\n", err)
	}
	fmt.Println(r)

	gameclient, err := dummyclient.Game(context.Background())

	go func() {
		for {
			err := gameclient.Send(&pb.Request{
				Sender: r,
				Action: "ping",
			})
			if err != nil {
				log.Fatalf("sending msg error :%s", err)
			}
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			in, err := gameclient.Recv()
			if err != nil {
				fmt.Printf("client receive error : %s\n", err)
				break
			}
			fmt.Printf("message received: %s %v\n", in.Status, in.Results)
		}
	}()

	inter := make(chan struct{}, 1)
	<-inter
}
