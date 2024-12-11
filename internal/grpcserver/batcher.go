package grpcserver

import (
	"context"
	"errors"
	"log"

	"github.com/GlebZigert/url_shortener.git/internal/packerr"
	pb "github.com/GlebZigert/url_shortener.git/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Batch struct {
	CorrelationID string
	URL           string
}

func (s *UrlShortenerServer) Batcher(ctx context.Context, in *pb.BatcherRequest) (*pb.BatcherResponce, error) {
	var resp pb.BatcherResponce
	var err error

	log.Println("grpc Batcher endpoint")

	//из входных собираем в структуру

	for _, itm := range in.Items {
		log.Println(itm.CorrelationId, " ", itm.OriginalUrl)

		short, err := s.service.Short(itm.OriginalUrl, -1)
		var conflict *packerr.ErrConflict409
		if err == nil || errors.As(err, &conflict) {

			var nested pb.BatcherResponce_Nested
			nested.CorrelationId = itm.CorrelationId
			nested.ShortUrl = short

			resp.Items = append(resp.Items, &nested)

		} else {
			return nil, status.Error(codes.Internal, "")
		}

	}

	return &resp, err
}
