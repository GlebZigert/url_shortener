package grpcserver

import (
	"context"
	"log"

	pb "github.com/GlebZigert/url_shortener.git/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UrlShortenerServer) GetURLs(ctx context.Context, in *pb.GetURLsRequest) (*pb.GetURLsResponce, error) {
	var resp pb.GetURLsResponce
	var err error
	log.Println("grpc GetURLs")

	user, ok := s.auc.CheckUID(ctx)
	if !ok {
		s.logger.Error("Нет данных о пользователе: ", map[string]interface{}{})

		return nil, status.Error(codes.Internal, "")
	}

	log.Println("uid ", user)

	for _, sh := range *s.service.GetAll() {
		if sh.UUID == int(user) {
			var itm pb.GetURLsResponce_Nested
			itm.Origin = sh.OriginalURL
			itm.Short = s.cfg.GetBaseURL() + "/" + sh.ShortURL
			resp.Items = append(resp.Items, &itm)
		}
	}

	if len(resp.Items) == 0 {
		return nil, status.Error(codes.NotFound, "")
	}

	return &resp, err
}
