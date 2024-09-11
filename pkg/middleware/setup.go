package middleware

import (
	"github.com/gin-gonic/gin"
	"go-metrics-demo/pkg/metrics"
	"go-metrics-demo/server/options"
	"strings"
)

func InstallMiddleware(opts *options.Options, ginEngine *gin.RouterGroup) {
	ginEngine.Use(Logger(), Cors(), Limiter(), Recovery(true), Validator())
	// monitor middleware
	monitor := metrics.NewMonitor(metrics.MonitorOptions{
		Namespace:     "tencent",
		SubSystemName: "metrics-demo",
		SkipPaths:     []string{strings.Join([]string{"/api/v1", "healthz"}, "/")},
	})
	monitor.UsedBy(opts.GinEngine)
}
