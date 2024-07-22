//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/lazywoo/mercury/feed/grpc"
	"github.com/lazywoo/mercury/feed/ioc"
	"github.com/lazywoo/mercury/pkg/wego"
)

var thirdProviderSet = wire.NewSet(
	ioc.InitLogger,
)

var serviceProviderSet = wire.NewSet()

func InitAPP() *wego.App {
	wire.Build(
		thirdProviderSet,
		serviceProviderSet,

		grpc.NewFeedEventServiceServer,
		ioc.InitGRPCxServer,
		wire.Struct(new(wego.App), "GRPCServer"),
	)
	return new(wego.App)
}
