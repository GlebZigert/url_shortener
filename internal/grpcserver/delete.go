package grpcserver

import (
	"context"
	"log"

	pb "github.com/GlebZigert/url_shortener.git/proto"
)

func (s *UrlShortenerServer) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponce, error) {
	var resp pb.DeleteResponce
	var err error

	log.Println("grpc Delete endpoint")

	for _, v := range in.Todel {
		log.Println(v)
	}

	return &resp, err
}
