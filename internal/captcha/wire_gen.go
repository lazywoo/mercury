// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"

	"github.com/lazywoo/mercury/internal/captcha/grpc"
	"github.com/lazywoo/mercury/internal/captcha/ioc"
	"github.com/lazywoo/mercury/internal/captcha/repository"
	"github.com/lazywoo/mercury/internal/captcha/repository/cache"
	"github.com/lazywoo/mercury/internal/captcha/service"
	"github.com/lazywoo/mercury/pkg/app"
)

// Injectors from wire.go:

func InitAPP() *app.App {
	cmdable := ioc.InitRedis()
	captchaCache := cache.NewCaptchaRedisCache(cmdable)
	captchaRepository := repository.NewCachedCaptchaRepository(captchaCache)
	client := ioc.InitEtcdClient()
	smsServiceClient := ioc.InitSmsServiceClient(client)
	captchaService := service.NewCaptchaService(captchaRepository, smsServiceClient)
	captchaServiceServer := grpc.NewCaptchaServiceServer(captchaService)
	logger := ioc.InitLogger()
	server := ioc.InitGRPCxServer(captchaServiceServer, logger)
	appApp := &app.App{
		GRPCServer: server,
	}
	return appApp
}

// wire.go:

var thirdProviderSet = wire.NewSet(ioc.InitLogger, ioc.InitRedis, ioc.InitEtcdClient, ioc.InitSmsServiceClient)

var svcProviderSet = wire.NewSet(grpc.NewCaptchaServiceServer, service.NewCaptchaService, repository.NewCachedCaptchaRepository, cache.NewCaptchaRedisCache)
