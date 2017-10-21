package main

import (
	"log"
	"net"

	"github.com/keigodasu/grpc-ref-error-handle/search"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type SearchServer struct{}

func (s *SearchServer) Search(ctx context.Context, in *search.SearchRequest) (*search.SearchResponse, error) {
	return &search.SearchResponse{}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	search.RegisterSearchServiceServer(s, &SearchServer{})
	s.Serve(l)
}
