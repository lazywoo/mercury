// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package startup

import (
	"github.com/google/wire"
	"github.com/lazywoo/mercury/internal/interactive/grpc"
	"github.com/lazywoo/mercury/internal/interactive/repository"
	"github.com/lazywoo/mercury/internal/interactive/repository/cache"
	"github.com/lazywoo/mercury/internal/interactive/repository/dao"
	"github.com/lazywoo/mercury/internal/interactive/service"
)

// Injectors from wire.go:

func InitInteractiveService() service.InteractiveService {
	gormDB := InitTestDB()
	interactiveDAO := dao.NewGORMInteractiveDAO(gormDB)
	cmdable := InitRedis()
	interactiveCache := cache.NewRedisInteractiveCache(cmdable)
	logger := InitLog()
	interactiveRepository := repository.NewCachedInteractiveRepository(interactiveDAO, interactiveCache, logger)
	interactiveService := service.NewInteractiveService(interactiveRepository, logger)
	return interactiveService
}

func InitInteractiveGRPCServer() *grpc.InteractiveServiceServer {
	gormDB := InitTestDB()
	interactiveDAO := dao.NewGORMInteractiveDAO(gormDB)
	cmdable := InitRedis()
	interactiveCache := cache.NewRedisInteractiveCache(cmdable)
	logger := InitLog()
	interactiveRepository := repository.NewCachedInteractiveRepository(interactiveDAO, interactiveCache, logger)
	interactiveService := service.NewInteractiveService(interactiveRepository, logger)
	interactiveServiceServer := grpc.NewInteractiveServiceServer(interactiveService)
	return interactiveServiceServer
}

// wire.go:

var thirdProvider = wire.NewSet(
	InitRedis,
	InitTestDB,
	InitLog,
	InitKafka,
)

var interactiveSvcProvider = wire.NewSet(service.NewInteractiveService, repository.NewCachedInteractiveRepository, dao.NewGORMInteractiveDAO, cache.NewRedisInteractiveCache)