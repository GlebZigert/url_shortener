package grpcserver

import (
	"context"

	pb "github.com/GlebZigert/url_shortener.git/proto"
)

func (s *UrlShortenerServer) Batcher(ctx context.Context, in *pb.BatcherRequest) (*pb.BatcherResponce, error) {
	var response pb.BatcherResponce
	var err error

	return &response, err
}
