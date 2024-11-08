// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"

	"github.com/lazywoo/mercury/internal/article/events"
	"github.com/lazywoo/mercury/internal/article/grpc"
	"github.com/lazywoo/mercury/internal/article/ioc"
	"github.com/lazywoo/mercury/internal/article/repository"
	"github.com/lazywoo/mercury/internal/article/repository/cache"
	"github.com/lazywoo/mercury/internal/article/repository/dao"
	"github.com/lazywoo/mercury/internal/article/service"
	"github.com/lazywoo/mercury/pkg/app"
)

// Injectors from wire.go:

func InitAPP() *app.App {
	logger := ioc.InitLogger()
	db := ioc.InitDB(logger)
	articleDAO := dao.NewGORMArticleDAO(db)
	cmdable := ioc.InitRedis()
	articleCache := cache.NewRedisArticleCache(cmdable)
	articleRepository := repository.NewCachedArticleRepository(articleDAO, articleCache, logger)
	client := ioc.InitEtcdClient()
	userServiceClient := ioc.InitUserRpcClient(client)
	saramaClient := ioc.InitKafka()
	syncProducer := ioc.NewSyncProducer(saramaClient)
	producer := events.NewSaramaSyncProducer(syncProducer)
	articleService := service.NewArticleService(articleRepository, userServiceClient, producer, logger)
	articleServiceServer := grpc.NewArticleServiceServer(articleService)
	server := ioc.InitGRPCxServer(articleServiceServer, logger)
	appApp := &app.App{
		GRPCServer: server,
	}
	return appApp
}

// wire.go:

var thirdProviderSet = wire.NewSet(ioc.InitLogger, ioc.InitDB, ioc.InitRedis, ioc.InitKafka, ioc.NewSyncProducer, ioc.InitEtcdClient, ioc.InitUserRpcClient)

var svcProviderSet = wire.NewSet(grpc.NewArticleServiceServer, events.NewSaramaSyncProducer, service.NewArticleService, repository.NewCachedArticleRepository, dao.NewGORMArticleDAO, cache.NewRedisArticleCache)
