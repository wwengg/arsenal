// @Title
// @Description
// @Author  Wangwengang  2021/8/18 下午10:06
// @Update  Wangwengang  2021/8/18 下午10:06
package config

import (
	"github.com/wwengg/arsenal/config/conf"
)

type Config struct {
	JWT           conf.JWT                   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`                     // jwt
	Zap           conf.Zap                   `mapstructure:"zap" json:"zap" yaml:"zap"`                     // logger
	TcpConfig     conf.TcpConfig             `mapstructure:"tcp-config" json:"tcpConfig" yaml:"tcp-config"` // logger
	Rpcx          conf.Rpcx                  `mapstructure:"rpcx" json:"rpcx" yaml:"rpcx"`
	EtcdV3        conf.EtcdV3                `mapstructure:"etcd-v3" json:"etcdV3" yaml:"etcd-v3"`
	Websocket     conf.Websocket             `mapstructure:"websocket" json:"websocket" yaml:"websocket"`
	Redis         conf.Redis                 `mapstructure:"redis" json:"redis" yaml:"redis"`
	RpcxRouterMap map[uint32]conf.RpcxRouter `mapstructure:"rpcx-router-map" json:"rpcxRouterMap" yaml:"rpcx-router-map"`

	// rainbow
	HttpGateway conf.HttpGateway `mapstructure:"http-gateway" json:"httpGateway" yaml:"http-gateway"`

	Mysql  conf.Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	DBList []conf.DB  `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
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
			Addr:     "127.0.0.1:8889",
			Network:  "tcp",
			Register: "etcdv3",
		},
		EtcdV3: conf.EtcdV3{Addr: []string{"127.0.0.1:23791", "127.0.0.1:23792", "127.0.0.1:23793"}},
		Websocket: conf.Websocket{
			Enable:           false,
			Addr:             "",
			ConnWriteTimeout: 0,
			ConnReadTimeout:  0,
			MaxHeaderLen:     0,
		},
		Redis: conf.Redis{
			Addr:       []string{"127.0.0.1:6379"},
			Db:         0,
			Password:   "",
			MasterName: "",
		},
		RpcxRouterMap: make(map[uint32]conf.RpcxRouter),
		HttpGateway: conf.HttpGateway{
			Env:  "dev",
			Addr: 8888,
		},
	}
}
