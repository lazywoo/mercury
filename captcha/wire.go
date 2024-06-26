//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/lazywoo/mercury/captcha/grpc"
	"github.com/lazywoo/mercury/captcha/ioc"
	"github.com/lazywoo/mercury/captcha/repository"
	"github.com/lazywoo/mercury/captcha/repository/cache"
	"github.com/lazywoo/mercury/captcha/service"
	"github.com/lazywoo/mercury/pkg/wego"
)

var thirdProviderSet = wire.NewSet(
	ioc.InitLogger,
	ioc.InitRedis,
	ioc.InitEtcdClient,
	ioc.InitSmsServiceClient,
)

var svcProviderSet = wire.NewSet(
	grpc.NewCaptchaServiceServer,
	service.NewCaptchaService,
	repository.NewCachedCaptchaRepository,
	cache.NewCaptchaRedisCache,
)

func InitAPP() *wego.App {
	wire.Build(
		thirdProviderSet,
		svcProviderSet,
		ioc.InitGRPCxServer,
		wire.Struct(new(wego.App), "GRPCServer"),
	)
	return new(wego.App)
}
