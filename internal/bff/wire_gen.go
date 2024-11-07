// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/lazywoo/mercury/internal/bff/ioc"
	"github.com/lazywoo/mercury/internal/bff/web"
	"github.com/lazywoo/mercury/internal/bff/web/jwt"
	"github.com/lazywoo/mercury/pkg/app"
)

// Injectors from wire.go:

func InitAPP() *app.App {
	cmdable := ioc.InitRedis()
	limiter := ioc.InitWebLimiter(cmdable)
	handler := jwt.NewRedisJWTHandler(cmdable)
	client := ioc.InitEtcdClient()
	userServiceClient := ioc.InitUserClient(client)
	captchaServiceClient := ioc.InitCaptchaClient(client)
	logger := ioc.InitLogger()
	userHandler := web.NewUserHandler(userServiceClient, captchaServiceClient, handler, logger)
	oauth2ServiceClient := ioc.InitOAuth2Client(client)
	wechatHandlerConfig := ioc.InitWechatHandlerConfig()
	oAuth2WechatHandler := web.NewOAuth2Handler(oauth2ServiceClient, userServiceClient, wechatHandlerConfig, handler)
	articleServiceClient := ioc.InitArticleClient(client)
	interactiveServiceClient := ioc.InitInteractiveClient(client)
	articleHandler := web.NewArticleHandler(articleServiceClient, interactiveServiceClient, logger)
	commentServiceClient := ioc.InitCommentClient(client)
	commentHandler := web.NewCommentHandler(commentServiceClient)
	server := ioc.InitWebServer(limiter, handler, userHandler, oAuth2WechatHandler, articleHandler, commentHandler, logger)
	appApp := &app.App{
		WebServer: server,
	}
	return appApp
}

// wire.go:

var thirdProviderSet = wire.NewSet(ioc.InitLogger, ioc.InitRedis, ioc.InitEtcdClient)

var hdlProviderSet = wire.NewSet(web.NewUserHandler, jwt.NewRedisJWTHandler, web.NewOAuth2Handler, web.NewArticleHandler, web.NewCommentHandler)

var cliProviderSet = wire.NewSet(ioc.InitUserClient, ioc.InitCaptchaClient, ioc.InitOAuth2Client, ioc.InitArticleClient, ioc.InitInteractiveClient, ioc.InitCommentClient)