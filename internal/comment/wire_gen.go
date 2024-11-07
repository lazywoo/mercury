// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"

	"github.com/lazywoo/mercury/internal/comment/grpc"
	"github.com/lazywoo/mercury/internal/comment/ioc"
	"github.com/lazywoo/mercury/internal/comment/repository"
	"github.com/lazywoo/mercury/internal/comment/repository/dao"
	"github.com/lazywoo/mercury/internal/comment/service"
)

// Injectors from wire.go:

func InitAPP() *App {
	logger := ioc.InitLogger()
	db := ioc.InitDB(logger)
	commentDAO := dao.NewCommentDAO(db)
	commentRepository := repository.NewCommentRepository(commentDAO, logger)
	commentService := service.NewCommentService(commentRepository)
	commentServiceServer := grpc.NewCommentServiceServer(commentService)
	server := ioc.InitGRPCxServer(commentServiceServer, logger)
	app := &App{
		server: server,
	}
	return app
}

// wire.go:

var thirdProviderSet = wire.NewSet(ioc.InitLogger, ioc.InitDB)

var serviceProviderSet = wire.NewSet(grpc.NewCommentServiceServer, service.NewCommentService, repository.NewCommentRepository, dao.NewCommentDAO)
