// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/lazywoo/mercury/internal/follow/grpc"
	"github.com/lazywoo/mercury/internal/follow/ioc"
	"github.com/lazywoo/mercury/internal/follow/repository"
	"github.com/lazywoo/mercury/internal/follow/repository/cache"
	"github.com/lazywoo/mercury/internal/follow/repository/dao"
	"github.com/lazywoo/mercury/internal/follow/service"
	"github.com/lazywoo/mercury/pkg/app"
)

// Injectors from wire.go:

func InitAPP() *app.App {
	logger := ioc.InitLogger()
	db := ioc.InitDB(logger)
	followDAO := dao.NewGORMFollowDAO(db)
	cmdable := ioc.InitRedis()
	followCache := cache.NewRedisFollowCache(cmdable)
	followRepository := repository.NewCachedFollowRepository(followDAO, followCache, logger)
	followService := service.NewFollowService(followRepository)
	followServiceServer := grpc.NewFollowServiceServer(followService)
	server := ioc.InitGRPCxServer(followServiceServer, logger)
	appApp := &app.App{
		GRPCServer: server,
	}
	return appApp
}

// wire.go:

var thirdProviderSet = wire.NewSet(ioc.InitDB, ioc.InitLogger, ioc.InitRedis)

var svcProviderSet = wire.NewSet(grpc.NewFollowServiceServer, service.NewFollowService, repository.NewCachedFollowRepository, dao.NewGORMFollowDAO, cache.NewRedisFollowCache)
