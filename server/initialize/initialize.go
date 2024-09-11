package initialize

import (
	"github.com/gin-gonic/gin"
	"go-metrics-demo/pkg/controller"
	"go-metrics-demo/pkg/logger"
	"go-metrics-demo/pkg/middleware"
	"go-metrics-demo/pkg/utils"
	"go-metrics-demo/server/options"
)

// InitServer init server
func InitServer(opts *options.Options) error {
	utils.PrintLogo()
	initGinEngine(opts)
	if err := initLogger(); err != nil {
		return err
	}
	InstallRouters(opts)
	return nil
}

// initLogger init logger
func initLogger() error {
	return logger.InitLogger()
}

// initGinEngine init default engine
func initGinEngine(opts *options.Options) {
	opts.GinEngine = gin.Default()
}

// InstallRouters install routers
func InstallRouters(opts *options.Options) {
	apiGroup := opts.GinEngine.Group("api/")
	middleware.InstallMiddleware(opts, apiGroup)
	controller.InstallRouter(apiGroup)
}
