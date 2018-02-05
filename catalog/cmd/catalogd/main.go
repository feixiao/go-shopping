package main

import (
	"github.com/feixiao/go-shopping/catalog/internal/platform/config"
	"github.com/feixiao/go-shopping/catalog/internal/platform/redis"
	"github.com/feixiao/go-shopping/catalog/internal/service"
	"github.com/feixiao/go-shopping/catalog/proto"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	svc := grpc.NewService(
		micro.Name(config.ServiceName),				// 微服务的名字
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),		// 重新注册的时间间隔
		micro.Version(config.Version),				// 微服务上报的版本信息
	)
	svc.Init()

	redisCatalogRepository := redis.NewRedisRepository(":6379")
	catalog.RegisterCatalogHandler(svc.Server(), service.NewCatalogService(redisCatalogRepository))

	if err := svc.Run(); err != nil {
		panic(err)
	}
}
