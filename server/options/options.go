package options

import (
	"github.com/gin-gonic/gin"
)

const (
	DefaultConfigFile string = "./config/config.yaml"
)

type Options struct {
	// GinEngine gin引擎对象
	GinEngine *gin.Engine
}

// NewOptions New one options with default ConfigFile
func NewOptions() (*Options, error) {
	return &Options{}, nil
}
