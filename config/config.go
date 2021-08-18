// @Title  
// @Description  
// @Author  Wangwengang  2021/8/18 下午10:06
// @Update  Wangwengang  2021/8/18 下午10:06
package config

import (
	"github.com/wwengg/arsenal/config/conf"
	"github.com/wwengg/arsenal/logger"
)

type Config struct {
	JWT conf.JWT `mapstructure:"jwt" json:"jwt" yaml:"jwt"` // jwt
	Zap conf.Zap `mapstructure:"zap" json:"zap" yaml:"zap"`  // logger
}

var ConfigHub *Config

func init(){
	ConfigHub = &Config{
		JWT: conf.JWT{
			SigningKey:  "wwengg",
			ExpiresTime: 604800,
			BufferTime:  86400,
		},
		Zap: conf.Zap{
			Level:         "info",
			Format:        "console",
			Prefix:        "[LOG]",
			Director:      "log",
			LinkName:      "latest_log",
			ShowLine:      true,
			EncodeLevel:   "LowercaseColorLevelEncoder",
			StacktraceKey: "stacktrace",
			LogInConsole:  true,
		},
	}

//	初始化全局日志
logger.ZapLog = logger.Zap()
}