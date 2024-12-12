package grpcserver

import (
	"context"
	"log"

	pb "github.com/GlebZigert/url_shortener.git/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UrlShortenerServer) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponce, error) {

	log.Println("grpc Delete endpoint")

	user, ok := s.auc.CheckUID(ctx)
	if !ok {
		s.logger.Error("Нет данных о пользователе: ", map[string]interface{}{})

		return nil, status.Error(codes.Internal, "")
	}

	for _, v := range in.Todel {
		log.Println(v)
	}

	go func() {
		cerr := s.service.Delete(in.Todel, user)
		if cerr != nil {
			s.logger.Error("Delete ", map[string]interface{}{
				"err": cerr.Error(),
			})

		}
	}()

	return nil, nil
}
