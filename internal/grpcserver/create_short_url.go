package grpcserver

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/GlebZigert/url_shortener.git/proto"
)

func (s *UrlShortenerServer) CreateShortURL(ctx context.Context, in *pb.CreateShortURLRequest) (*pb.CreateShortURLResponse, error) {
	var response pb.CreateShortURLResponse

	user, ok := s.auc.CheckUID(ctx)
	if !ok {
		s.logger.Error("Нет данных о пользователе: ", map[string]interface{}{})

		return nil, status.Error(codes.Internal, "")
	}

	res, err := s.service.Short(in.Origin, user)
	response.Short = res
	if err != nil {
		response.Error = err.Error()
	}

	return &response, nil
}
