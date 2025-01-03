// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/lazywoo/mercury/internal/interactive/events"
	"github.com/lazywoo/mercury/internal/interactive/grpc"
	"github.com/lazywoo/mercury/internal/interactive/ioc"
	"github.com/lazywoo/mercury/internal/interactive/repository"
	"github.com/lazywoo/mercury/internal/interactive/repository/cache"
	"github.com/lazywoo/mercury/internal/interactive/repository/dao"
	"github.com/lazywoo/mercury/internal/interactive/service"
	"github.com/lazywoo/mercury/pkg/app"
)

// Injectors from wire.go:

func InitAPP() *app.App {
	logger := ioc.InitLogger()
	srcDB := ioc.InitSrcDB(logger)
	dstDB := ioc.InitDstDB(logger)
	dualWritePool := ioc.InitDualWritePool(srcDB, dstDB)
	db := ioc.InitDualWriteDB(dualWritePool)
	interactiveDAO := dao.NewGORMInteractiveDAO(db)
	cmdable := ioc.InitRedis()
	interactiveCache := cache.NewRedisInteractiveCache(cmdable)
	interactiveRepository := repository.NewCachedInteractiveRepository(interactiveDAO, interactiveCache, logger)
	interactiveService := service.NewInteractiveService(interactiveRepository, logger)
	interactiveServiceServer := grpc.NewInteractiveServiceServer(interactiveService)
	server := ioc.InitGRPCxServer(interactiveServiceServer, logger)
	client := ioc.InitKafka()
	interactiveReadEventConsumer := events.NewInteractiveReadEventConsumer(client, interactiveRepository, logger)
	consumer := ioc.InitFixDataConsumer(srcDB, dstDB, client, logger)
	v := ioc.NewConsumers(interactiveReadEventConsumer, consumer)
	appApp := &app.App{
		GRPCServer: server,
		Consumers:  v,
	}
	return appApp
}

// wire.go:

var thirdProvider = wire.NewSet(ioc.InitSrcDB, ioc.InitDstDB, ioc.InitDualWritePool, ioc.InitDualWriteDB, ioc.InitRedis, ioc.InitKafka, ioc.InitLogger, ioc.NewSyncProducer)

var interactiveSvcProvider = wire.NewSet(service.NewInteractiveService, repository.NewCachedInteractiveRepository, dao.NewGORMInteractiveDAO, cache.NewRedisInteractiveCache)

var migratorSet = wire.NewSet(ioc.InitMigratorProducer, ioc.InitFixDataConsumer, ioc.InitMigratorWeb)
