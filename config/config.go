package config

type LogOptions struct {
	// Level 日志级别
	Level string `mapstructure:"level"`
	// Filename 日志文件位置
	Filename string `mapstructure:"fileName"`
	// MaxSize 日志文件最大大小(MB)
	MaxSize int `mapstructure:"maxSize"`
	// MaxAge 保留旧日志文件的最大天数
	MaxAge int `mapstructure:"maxAge"`
	// MaxBackups 最大保留日志个数
	MaxBackups int `mapstructure:"maxBackups"`
}
