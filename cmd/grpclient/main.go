package main

import (
	"context"
	"log"

	pb "github.com/GlebZigert/url_shortener.git/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial(":3200", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewUrlShortenerClient(conn)

	Test(client)

}

func Test(c pb.UrlShortenerClient) {
	req := &pb.CreateShortURLRequest{Origin: "www.leningradspb.ru"}
	ctx := context.Background()
	var header, trailer metadata.MD
	_, err := c.CreateShortURL(ctx,
		req,
		grpc.Header(&header),   // will retrieve header
		grpc.Trailer(&trailer)) // will retrieve trailer)
	if err != nil {
		log.Println(err.Error())

	}

	authHeader := header.Get("authorisation")
	log.Println(":::", authHeader)

	jwt := authHeader[0]
	log.Println(":::", jwt)
	md := metadata.Pairs("authorisation", jwt)
	ctx = metadata.NewOutgoingContext(context.Background(), md)
	req = &pb.CreateShortURLRequest{Origin: "http.yandex.ru"}

	_, err = c.CreateShortURL(ctx,
		req,
		grpc.Header(&header), // will retrieve header
		grpc.Trailer(&trailer))

	ctx = context.Background()
	req = &pb.CreateShortURLRequest{Origin: "http.google.com"}

	resp, err := c.CreateShortURL(ctx,
		req,
		grpc.Header(&header), // will retrieve header
		grpc.Trailer(&trailer))

	if err != nil {
		log.Println(err.Error())
		return

	}
	short := resp.Short

	ctx = metadata.NewOutgoingContext(context.Background(), md)
	getUrlreq := &pb.GetURLRequest{Short: short}

	getUrlresp, err := c.GetURL(ctx,
		getUrlreq,
		grpc.Header(&header), // will retrieve header
		grpc.Trailer(&trailer))

	if err != nil {
		log.Println(err.Error())
		return

	}

	log.Println(getUrlresp.Origin)

	br := &pb.BatcherRequest{Items: []*pb.BatcherRequest_Nested{
		&pb.BatcherRequest_Nested{CorrelationId: "1", OriginalUrl: "url0001"},
		&pb.BatcherRequest_Nested{CorrelationId: "2", OriginalUrl: "url0002"},
	}}

	batchresp, err := c.Batcher(context.Background(), br)

	if err != nil {
		log.Println(err.Error())
		return

	}
	getUrlreq = &pb.GetURLRequest{Short: batchresp.Items[0].ShortUrl}
	getUrlresp, err = c.GetURL(ctx,
		getUrlreq,
		grpc.Header(&header), // will retrieve header
		grpc.Trailer(&trailer))

	if err != nil {
		log.Println(err.Error())
		return

	}

	log.Println(getUrlresp)
	ctx = metadata.NewOutgoingContext(context.Background(), md)
	urlsreq := &pb.GetURLsRequest{}

	geturlresp, err := c.GetURLs(ctx,
		urlsreq,
		grpc.Header(&header), // will retrieve header
		grpc.Trailer(&trailer))

	for _, v := range geturlresp.Items {
		log.Println(v.Origin, " ", v.Short)
	}

}
