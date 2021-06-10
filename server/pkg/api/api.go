package api

import (
	"fmt"

	pb "github.com/PutskouDzmitry/golang-trainnig-final-task/proto/go_proto"

	"github.com/sirupsen/logrus"
	"gopkg.in/mcuadros/go-syslog.v2"
)

type EventServer struct {
	channel syslog.LogPartsChannel
	server  *syslog.Server
}

func (u EventServer) GetEvent(stream pb.Service_GetEventServer) error {
	logrus.Info("Client connected")
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)
	u.CheckClient(stream, handler)
	u.server.SetHandler(handler)
	eventChannel := make(chan *pb.Event)
	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			event := &pb.Event{
				Facility: fmt.Sprintf("%v", logParts["facility"]),
				Message:  fmt.Sprintf("%v", logParts["message"]),
				Severity: fmt.Sprintf("%v", logParts["severity"]),
				Time:     fmt.Sprintf("%v", logParts["timestamp"]),
			}
			eventChannel <- event
		}
	}(channel)
	for {
		err := SendInfo(stream, eventChannel)
		if err != nil {
			logrus.Fatal("got an error form client ", err)
		}
	}
}

func SendInfo(stream pb.Service_GetEventServer, eventChannel <-chan *pb.Event) error {
	for {
		if err := stream.Send(&pb.EventResponse{Event: <-eventChannel}); err != nil {
			return err
		}
	}
}

func (u EventServer) CheckClient(stream pb.Service_GetEventServer, handler *syslog.ChannelHandler) {
	go func() {
		for {
			_, err := stream.Recv()
			if err != nil {
				logrus.Info("Client disconnected")
				handler.SetChannel(u.channel)
				return
			}
		}
	}()
}

func NewEventServer(channel syslog.LogPartsChannel, server *syslog.Server) *EventServer {
	return &EventServer{
		channel: channel,
		server:  server,
	}
}
