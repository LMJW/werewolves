package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"

	pb "github.com/LMJW/werewolves/server/proto"
)

const (
	port = ":50051"
)

// server is used to implement Werewolf.WerewolfServer.
type server struct {
	count     int
	broadcast map[int]chan int
	state     bool
	mux       sync.Mutex
}

func (s *server) Register(l chan int) {
	s.mux.Lock()
	s.broadcast[len(s.broadcast)] = l
	s.mux.Unlock()
	if s.state == true {
		return
	}
	s.state = true
	s.Start()
}

func (s *server) remove(i int) {
	s.mux.Lock()
	defer s.mux.Unlock()
	delete(s.broadcast, i)
	log.Printf("Player (%d) removed for server broadcast list\n", i)
}

func (s *server) Start() {
	go func() {
		for {
			time.Sleep(time.Second)
			s.count++
			for k, v := range s.broadcast {
				select {
				case v <- s.count:
				default:
					fmt.Printf("player (%d) is blocking\n", k)
					s.remove(k)
				}
			}
		}
	}()
}

// CreateID implements Werewolf.WerewolfServer.
func (s *server) CreateID(ctx context.Context, in *pb.Signature) (*pb.Player, error) {
	log.Printf("Received: %v\n", in.Displayname)
	return &pb.Player{}, nil
}

// Game implements Werewolf.WerewolfServer
func (s *server) Game(srv pb.Werewolf_GameServer) error {
	var ll = make(chan int)
	s.Register(ll)
	go func() {
		for v := range ll {
			num := strconv.Itoa(v)
			err := srv.Send(&pb.Response{
				Status:  "tick",
				Results: []string{num},
			})
			if err != nil {
				fmt.Printf("error sending message: %s\n", err)
				return
			}
		}
	}()

	for {
		in, err := srv.Recv()
		if err != nil {
			return err
		}
		fmt.Printf("Server receive: %s\n", in.Action)
		err = srv.Send(&pb.Response{
			Status: "pong",
		})
		if err != nil {
			fmt.Printf("error sending message: %s\n", err)
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWerewolfServer(s, &server{
		count:     0,
		broadcast: make(map[int]chan int),
		state:     false,
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
