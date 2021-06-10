package main

import (
	"log"
	"net"

	pb "github.com/PutskouDzmitry/golang-trainnig-final-task/proto/go_proto"
	"github.com/PutskouDzmitry/golang-trainnig-final-task/server/pkg/api"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/mcuadros/go-syslog.v2"
)

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	serverGrpc := grpc.NewServer()
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)
	server, err := ConfServer(handler)
	if err != nil {
		logrus.Fatal(err)
	}
	OutputInfoInConsole(channel)
	pb.RegisterServiceServer(serverGrpc, api.NewEventServer(channel, server))
	if err = serverGrpc.Serve(listener); err != nil {
		log.Fatal(err)
	}
	server.Wait()
}

func ConfServer(handler *syslog.ChannelHandler) (*syslog.Server, error) {
	server := syslog.NewServer()
	server.SetFormat(syslog.RFC5424)
	server.SetHandler(handler)
	err := server.ListenUDP("0.0.0.0:1514")
	if err != nil {
		return nil, err
	}
	err = server.Boot()
	if err != nil {
		return nil, err
	}
	return server, nil
}

func OutputInfoInConsole(channel syslog.LogPartsChannel) {
	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			logrus.WithFields(logrus.Fields{
				"facility": logParts["facility"],
				"msg":      logParts["message"],
				"severity": logParts["severity"],
				"time":     logParts["timestamp"],
			}).Info()
		}
	}(channel)
}
