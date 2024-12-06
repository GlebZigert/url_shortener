package main

import (
	"context"
	"log"

	pb "github.com/GlebZigert/url_shortener.git/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	_, err := c.CreateShortURL(ctx, req)
	if err != nil {
		log.Println(err.Error())
	}
}
