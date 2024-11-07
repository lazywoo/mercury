// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/lazywoo/mercury/internal/ranking/grpc"
	"github.com/lazywoo/mercury/internal/ranking/ioc"
	"github.com/lazywoo/mercury/internal/ranking/repository"
	"github.com/lazywoo/mercury/internal/ranking/repository/cache"
	"github.com/lazywoo/mercury/internal/ranking/service"
	"github.com/lazywoo/mercury/pkg/app"
)

// Injectors from wire.go:

func InitAPP() *app.App {
	client := ioc.InitEtcdClient()
	articleServiceClient := ioc.InitArticleRpcClient(client)
	interactiveServiceClient := ioc.InitInteractiveRpcClient(client)
	cmdable := ioc.InitRedis()
	rankingRedisCache := cache.NewRankingRedisCache(cmdable)
	rankingLocalCache := cache.NewRankingLocalCache()
	rankingRepository := repository.NewRankingCachedRepository(rankingRedisCache, rankingLocalCache)
	rankingService := service.NewBatchRankingService(articleServiceClient, interactiveServiceClient, rankingRepository)
	rankingServiceServer := grpc.NewRankingServiceServer(rankingService)
	logger := ioc.InitLogger()
	server := ioc.InitGRPCxServer(rankingServiceServer, logger)
	rlockClient := ioc.InitRLockClient(cmdable)
	rankingJob := ioc.InitRankingJob(rankingService, rlockClient, logger)
	cron := ioc.InitTasks(logger, rankingJob)
	appApp := &app.App{
		GRPCServer: server,
		Cron:       cron,
	}
	return appApp
}

// wire.go:

var thirdProviderSet = wire.NewSet(ioc.InitLogger, ioc.InitRedis, ioc.InitEtcdClient, ioc.InitArticleRpcClient, ioc.InitInteractiveRpcClient)

var svcProviderSet = wire.NewSet(service.NewBatchRankingService, repository.NewRankingCachedRepository, cache.NewRankingLocalCache, cache.NewRankingRedisCache)

var cronProviderSet = wire.NewSet(ioc.InitTasks, ioc.InitRankingJob, ioc.InitRLockClient)
