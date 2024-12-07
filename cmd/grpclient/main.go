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

	md := metadata.Pairs("authorisation", jwt)
	ctx = metadata.NewOutgoingContext(context.Background(), md)
	req = &pb.CreateShortURLRequest{Origin: "http.yandex.ru"}

	_, err = c.CreateShortURL(ctx,
		req,
		grpc.Header(&header), // will retrieve header
		grpc.Trailer(&trailer))

	ctx = context.Background()
	req = &pb.CreateShortURLRequest{Origin: "http.google.com"}

	_, err = c.CreateShortURL(ctx,
		req,
		grpc.Header(&header), // will retrieve header
		grpc.Trailer(&trailer))

}
