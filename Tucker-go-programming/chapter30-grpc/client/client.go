package main

import (
	"bufio"
	"context"
	"flag"
	"io"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpcchat/chatproto"
)

var id = flag.String("id", "1", "The id name")
var serverAdddr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")

func main() {
	flag.Parse()
	conn, err := grpc.Dial(
		*serverAdddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("fail to dial : %v", err)
	}

	defer conn.Close()

	client := pb.NewChatServiceClient(conn)
	runChat(client)
}

func runChat(client pb.ChatServiceClient) {
	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatalf("open stream error: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			in, err := stream.Recv()
			if err != io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note: %v", err)
			}
			log.Printf("Sender:%s Message:%s", in.Sender, in.Message)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		if strings.ToLower(msg) == "exit" {
			break
		}
		msg1 := &pb.ChatMsg{
			Sender:  *id,
			Message: scanner.Text(),
		}
		stream.Send(msg1)
	}
	stream.CloseSend()
	<-waitc
}
