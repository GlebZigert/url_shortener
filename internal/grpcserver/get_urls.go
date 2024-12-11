package grpcserver

import (
	"context"
	"log"

	pb "github.com/GlebZigert/url_shortener.git/proto"
)

func (s *UrlShortenerServer) GetURLs(ctx context.Context, in *pb.GetURLsRequest) (*pb.GetURLsResponce, error) {
	var response pb.GetURLsResponce
	var err error
	log.Println("grpc GetURLs")

	return &response, err
}
