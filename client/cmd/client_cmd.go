package main

import (
	"context"

	pb "github.com/PutskouDzmitry/golang-trainnig-final-task/proto/go_proto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logrus.Fatal(err)
	}
	defer conn.Close()
	stream, err := pb.NewServiceClient(conn).GetEvent(context.Background())
	if err != nil {
		logrus.Fatal(err)
	}
	for {
		event, err := SendInfo(stream)
		if err != nil {
			logrus.Fatal(err)
		}
		logrus.Info(event)
	}
}

func SendInfo(stream pb.Service_GetEventClient) (*pb.EventResponse, error) {
	for {
		event, err := stream.Recv()
		if err != nil {
			return &pb.EventResponse{}, err
		}
		err = stream.Send(&pb.EventRequest{})
		if err != nil {
			return &pb.EventResponse{}, err
		}
		return event, nil
	}
}
