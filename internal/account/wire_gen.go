// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/lazywoo/mercury/internal/account/grpc"
	"github.com/lazywoo/mercury/internal/account/ioc"
	"github.com/lazywoo/mercury/internal/account/repository"
	"github.com/lazywoo/mercury/internal/account/repository/dao"
	"github.com/lazywoo/mercury/internal/account/service"
	"github.com/lazywoo/mercury/pkg/app"
)

// Injectors from wire.go:

func InitAPP() *app.App {
	logger := ioc.InitLogger()
	db := ioc.InitDB(logger)
	accountDAO := dao.NewAccountDAO(db)
	accountRepository := repository.NewAccountRepository(accountDAO)
	accountService := service.NewAccountServiceServer(accountRepository)
	accountServiceServer := grpc.NewAccountServiceServer(accountService)
	server := ioc.InitGRPCxServer(accountServiceServer, logger)
	appApp := &app.App{
		GRPCServer: server,
	}
	return appApp
}

// wire.go:

var thirdProviderSet = wire.NewSet(ioc.InitDB, ioc.InitLogger)

var svcProviderSet = wire.NewSet(grpc.NewAccountServiceServer, service.NewAccountServiceServer, repository.NewAccountRepository, dao.NewAccountDAO)