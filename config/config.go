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
	JWT       conf.JWT       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`                     // jwt
	Zap       conf.Zap       `mapstructure:"zap" json:"zap" yaml:"zap"`                     // logger
	TcpConfig conf.TcpConfig `mapstructure:"tcp-config" json:"tcpConfig" yaml:"tcp-config"` // logger
	Rpcx      conf.Rpcx      `mapstructure:"rpcx" json:"rpcx" yaml:"rpcx"`
	EtcdV3    conf.EtcdV3    `mapstructure:"etcd-v3" json:"etcdV3" yaml:"etcd-v3"`
}

var ConfigHub *Config

func init() {
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
		TcpConfig: conf.TcpConfig{
			Name:             "pigeon",
			Ip:               "0.0.0.0",
			TcpPort:          8999,
			MaxConn:          3,
			WorkerPoolSize:   10,
			MaxWorkerTaskLen: 1024,
			MaxMsgChanLen:    1024,
			MaxPacketSize:    4096,
		},
		Rpcx: conf.Rpcx{
			BasePath: "rpcx",
			Addr:     "0.0.0.0:8889",
			Network:  "tcp",
			Register: "etcdv3",
		},
		EtcdV3: conf.EtcdV3{Addr: []string{"127.0.0.1:23791", "127.0.0.1:23792", "127.0.0.1:23793"}},
	}

	//	初始化全局日志
	logger.ZapLog = logger.Zap()
}
