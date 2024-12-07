package grpcserver

import (
	"context"

	pb "github.com/GlebZigert/url_shortener.git/proto"
)

func (s *UrlShortenerServer) GetURL(ctx context.Context, in *pb.GetURLRequest) (*pb.GetURLResponse, error) {
	var response pb.GetURLResponse

	return &response, nil
}
