// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"webook/internal/api"
	"webook/internal/repository"
	cache2 "webook/internal/repository/cache/captcha"
	"webook/internal/repository/cache/user"
	"webook/internal/repository/dao"
	"webook/internal/service"
	"webook/ioc"
)

// Injectors from wire.go:

func InitWebServer() *gin.Engine {
	cmdable := ioc.InitRedis()
	limiter := ioc.InitLimiter(cmdable)
	v := ioc.InitMiddlewares(limiter)
	db := ioc.InitDB()
	userDao := dao.NewUserGormDao(db)
	userCache := cache.NewUserRedisCache(cmdable)
	userRepository := repository.NewUserCachedRepository(userDao, userCache)
	userService := service.NewUserServiceV1(userRepository)
	captchaCache := cache2.NewCaptchaRedisCache(cmdable)
	captchaRepository := repository.NewCaptchaCachedRepository(captchaCache)
	smsService := ioc.InitSMSService()
	captchaService := service.NewCaptchaServiceV1(captchaRepository, smsService)
	userHandler := api.NewUserHandler(userService, captchaService)
	engine := ioc.InitWebServer(v, userHandler)
	return engine
}
