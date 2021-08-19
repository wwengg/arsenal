// @Title
// @Description
// @Author  Wangwengang  2021/8/19 下午5:22
// @Update  Wangwengang  2021/8/19 下午5:22
package rpcx

import (
	"github.com/rpcxio/rpcxplus/grpcx"
	"github.com/smallnest/rpcx/server"
	"pigeon-logic/service"

	"github.com/wwengg/arsenal/config"
	"github.com/wwengg/arsenal/sdk/rpcx/plugin"
)

type Server interface {
	RegisterName(name string, rcvr interface{}, metadata string)
	Serve()
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
	s.rpcxServer.Serve(config.ConfigHub.Rpcx.Network, config.ConfigHub.Rpcx.Addr)
}
