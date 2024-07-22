package grpc

import (
	"context"
	feedv1 "github.com/lazywoo/mercury/api/proto/gen/feed/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FeedEventServiceServer struct {
	feedv1.UnimplementedFeedServiceServer
}

func NewFeedEventServiceServer() *FeedEventServiceServer {
	return &FeedEventServiceServer{}
}

func (srv *FeedEventServiceServer) CreateFeedEvent(context.Context, *feedv1.CreateFeedEventRequest) (*feedv1.CreateFeedEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFeedEvent not implemented")
}

func (srv *FeedEventServiceServer) FindFeedEvents(context.Context, *feedv1.FindFeedEventsRequest) (*feedv1.FindFeedEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindFeedEvents not implemented")
}

func (srv *FeedEventServiceServer) Register(registrar grpc.ServiceRegistrar) {
	feedv1.RegisterFeedServiceServer(registrar, srv)
}
