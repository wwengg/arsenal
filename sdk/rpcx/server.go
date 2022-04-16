// @Title
// @Description
// @Author  Wangwengang  2021/8/19 下午5:22
// @Update  Wangwengang  2021/8/19 下午5:22
package rpcx

import (
	"github.com/smallnest/rpcx/server"
	"go.uber.org/zap"

	"github.com/wwengg/arsenal/config"
	"github.com/wwengg/arsenal/logger"
	"github.com/wwengg/arsenal/sdk/rpcx/plugin"
)

type Server interface {
	RegisterName(name string, rcvr interface{}, metadata string)
	Serve()
	GetServer() *server.Server
}

type RpcxServer struct {
	rpcxServer *server.Server
}

func NewRpcxServer() Server {
	s := RpcxServer{rpcxServer: server.NewServer()}
	// 添加etcdv4注册中心
	plugin.AddRegisteryPlugin(s.rpcxServer)

	return &s
}

func (s *RpcxServer) RegisterName(name string, rcvr interface{}, metadata string) {
	s.rpcxServer.RegisterName(name, rcvr, metadata)
}

func (s *RpcxServer) Serve() {
	logger.ZapLog.Info("rpcxServer Start", zap.String("addr", config.ConfigHub.Rpcx.Addr))
	s.rpcxServer.Serve(config.ConfigHub.Rpcx.Network, config.ConfigHub.Rpcx.Addr)
}

func (s *RpcxServer) GetServer() *server.Server {
	return s.rpcxServer
}