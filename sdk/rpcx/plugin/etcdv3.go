// @Title
// @Description
// @Author  Wangwengang  2021/8/19 上午10:59
// @Update  Wangwengang  2021/8/19 上午10:59
package plugin

import (
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/rpcxio/rpcx-etcd/serverplugin"
	"github.com/smallnest/rpcx/server"
	"go.uber.org/zap"

	"github.com/wwengg/arsenal/config"
	"github.com/wwengg/arsenal/logger"
)

func AddRegisteryPlugin(s *server.Server) {
	r := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: "tcp@" + config.ConfigHub.Rpcx.Addr,
		EtcdServers:    config.ConfigHub.EtcdV3.Addr,
		BasePath:       config.ConfigHub.Rpcx.BasePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		logger.ZapLog.Error(err.Error())
	}
	s.Plugins.Add(r)
	logger.ZapLog.Info("register etcdv3 start", zap.Any("ServiceAddress", r.ServiceAddress), zap.Any("etcdServers", r.EtcdServers))
}
