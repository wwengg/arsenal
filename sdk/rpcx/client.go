// @Title
// @Description
// @Author  Wangwengang  2021/8/19 下午4:58
// @Update  Wangwengang  2021/8/19 下午4:58
package rpcx

import (
	"fmt"
	"sync"

	etcdclient "github.com/rpcxio/rpcx-etcd/client"
	"github.com/smallnest/rpcx/client"

	"github.com/wwengg/arsenal/config"
	"github.com/wwengg/arsenal/logger"
)

type RpcxClients struct {
	FailMode   client.FailMode
	SelectMode client.SelectMode
	Option     client.Option

	mu               sync.RWMutex
	xclients         map[string]client.XClient
	ServiceDiscovery client.ServiceDiscovery

	seq uint64
}

var (
	RpcxClientsObj *RpcxClients
)

func init() {
	RpcxClientsObj = &RpcxClients{
		FailMode:   client.FailMode(int(client.Failover)),
		SelectMode: client.SelectMode(int(client.RoundRobin)),
		Option:     client.DefaultOption,
		xclients:   make(map[string]client.XClient),
	}
}

func (rc *RpcxClients) SetFailMode(mode client.FailMode){
	rc.FailMode = mode
}

func (rc *RpcxClients) SetSelectMode(selectMode client.SelectMode){
	rc.SelectMode = selectMode
}

func (rc *RpcxClients) SetOption(option client.Option){
	rc.Option = option
}

func (rc *RpcxClients) SetupServiceDiscovery() {
	var err error
	rc.ServiceDiscovery, err = etcdclient.NewEtcdV3DiscoveryTemplate("yf", config.ConfigHub.EtcdV3.Addr, true, nil)
	if err != nil {
		logger.ZapLog.Panic(err.Error())
		panic(err)
	}
}

func (rc *RpcxClients) GetXClient(servicePath string) (xc client.XClient, err error) {
	defer func() {
		if e := recover(); e != nil {
			if ee, ok := e.(error); ok {
				err = ee
				return
			}

			err = fmt.Errorf("failed to get xclient: %v", e)
		}
	}()
	rc.mu.RLock()
	if rc.xclients[servicePath] == nil {
		d, err := rc.ServiceDiscovery.Clone(servicePath)
		if err != nil {
			return nil, err
		}
		rc.xclients[servicePath] = client.NewXClient(servicePath, rc.FailMode, rc.SelectMode, d, rc.Option)
	}
	xc = rc.xclients[servicePath]
	rc.mu.RUnlock()

	return xc, err
}

func (rc *RpcxClients) NewXClient(servicePath string) (xc client.XClient, err error) {
	d, err := rc.ServiceDiscovery.Clone(servicePath)
	if err != nil {
		return nil, err
	}
	xc = client.NewXClient(servicePath, rc.FailMode, rc.SelectMode, d, rc.Option)

	return xc, err
}
