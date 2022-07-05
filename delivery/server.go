package delivery

import (
	"go-api-with-gin2/config"
	"go-api-with-gin2/delivery/controller"
	"go-api-with-gin2/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
}

func Server() *appServer {
	r := gin.Default()
	appConfig := config.NewConfig()
	infra := manager.NewInfra(&appConfig)
	repoManager := manager.NewRepoManager(infra)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	host := appConfig.Url
	return &appServer{
		useCaseManager: useCaseManager,
		engine:         r,
		host:           host,
	}
}

func (a *appServer) initControllers() {
	controller.NewProductController(a.engine, a.useCaseManager.CreateProductUseCase(), a.useCaseManager.ListProductUseCase())
}

func (a *appServer) Run() {
	a.initControllers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
