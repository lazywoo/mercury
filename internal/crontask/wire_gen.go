// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"

	"github.com/lazywoo/mercury/internal/crontask/grpc"
	"github.com/lazywoo/mercury/internal/crontask/ioc"
	"github.com/lazywoo/mercury/internal/crontask/repository"
	"github.com/lazywoo/mercury/internal/crontask/repository/dao"
	"github.com/lazywoo/mercury/internal/crontask/service"
	"github.com/lazywoo/mercury/pkg/app"
)

// Injectors from wire.go:

func InitAPP() *app.App {
	logger := ioc.InitLogger()
	db := ioc.InitDB(logger)
	taskDAO := dao.NewGORMTaskDAO(db)
	taskRepository := repository.NewPreemptTaskRepository(taskDAO)
	taskService := service.NewTaskService(taskRepository, logger)
	cronJobServiceServer := grpc.NewCronJobServiceServer(taskService)
	server := ioc.InitGRPCxServer(cronJobServiceServer, logger)
	appApp := &app.App{
		GRPCServer: server,
	}
	return appApp
}

// wire.go:

var thirdProviderSet = wire.NewSet(ioc.InitDB, ioc.InitLogger)

var svcProviderSet = wire.NewSet(grpc.NewCronJobServiceServer, service.NewTaskService, repository.NewPreemptTaskRepository, dao.NewGORMTaskDAO)
