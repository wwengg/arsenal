// @Title
// @Description
// @Author  Wangwengang  2021/8/17 下午5:08
// @Update  Wangwengang  2021/8/17 下午5:08
package impl

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/wwengg/proto/identity"
	"go.uber.org/zap"

	"github.com/wwengg/arsenal/anet"
	"github.com/wwengg/arsenal/config"
	"github.com/wwengg/arsenal/logger"
	"github.com/wwengg/arsenal/sdk/rpcx"
)

//Server 接口实现，定义一个Server服务类
type Server struct {
	//服务器的名称
	Name string
	//tcp4 or other
	IPVersion string
	//服务绑定的IP地址
	IP string
	//服务绑定的端口
	Port int
	// Websocket Addr
	WsAddr string
	//当前Server的消息管理模块，用来绑定MsgID和对应的处理方法
	msgHandler anet.MsgHandle
	//当前Server的链接管理器
	ConnMgr anet.ConnManager
	//该Server的连接创建时Hook函数
	OnConnStart func(conn anet.Connection)
	//该Server的连接断开时的Hook函数
	OnConnStop func(conn anet.Connection)

	packet anet.Packet

	identityClient *identity.IdentityClient
}

//NewServer 创建一个服务器句柄
func NewServer(opts ...Option) anet.Server {
	//printLogo()
	rpcx.RpcxClientsObj.SetupServiceDiscovery()
	xclient, err := rpcx.RpcxClientsObj.GetXClient("Identity")
	if err != nil {
		logger.ZapLog.Error("Identity service not found")
	}

	s := &Server{
		Name:           config.ConfigHub.TcpConfig.Name,
		IPVersion:      "tcp4",
		IP:             config.ConfigHub.TcpConfig.Ip,
		Port:           config.ConfigHub.TcpConfig.TcpPort,
		WsAddr:         config.ConfigHub.Websocket.Addr,
		msgHandler:     NewMsgHandle(),
		ConnMgr:        NewConnManager(),
		packet:         NewDataPack(),
		identityClient: identity.NewIdentityClient(xclient),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Server) GenID() uint64 {
	reply, err := s.identityClient.GetId(context.Background(), nil)
	if err != nil {
		return 0
	}
	return uint64(reply.Id)
}

//============== 实现 anet.Server 里的全部接口方法 ========

//Start 开启Tcp网络服务
func (s *Server) StartTcp() {
	fmt.Printf("[START] Server name: %s,listenner at IP: %s, Port %d is starting\n", s.Name, s.IP, s.Port)

	//开启一个go去做服务端Linster业务
	go func() {
		//0 启动worker工作池机制
		s.msgHandler.StartWorkerPool()

		//1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		//2 监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			panic(err)
		}

		//已经监听成功
		fmt.Println("start Zinx server  ", s.Name, " succ, now listenning...")

		//TODO server.go 应该有一个自动生成ID的方法

		//3 启动server网络连接业务
		for {
			//3.1 阻塞等待客户端建立连接请求
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}
			fmt.Println("Get conn remote addr = ", conn.RemoteAddr().String())

			//3.2 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
			if s.ConnMgr.Len() >= 100 {
				conn.Close()
				continue
			}

			//3.3 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
			dealConn := NewConnection(s, conn, s.GenID(), s.msgHandler)

			//3.4 启动当前链接的处理业务
			go dealConn.Start()
		}
	}()
}

//Start Websocket网络服务
func (s *Server) StartWebsocket() {
	logger.ZapLog.Info("Start Websocket server", zap.String("addr", s.WsAddr))
	httpServer := &http.Server{
		Addr: s.WsAddr,
		Handler: &WsHandler{upgrader: websocket.Upgrader{
			HandshakeTimeout: 0,
			ReadBufferSize:   0,
			WriteBufferSize:  0,
			WriteBufferPool:  nil,
			Subprotocols:     nil,
			Error:            nil,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			EnableCompression: false,
		},
			sv: s},
		ReadTimeout:    time.Second * time.Duration(config.ConfigHub.Websocket.ConnReadTimeout),
		WriteTimeout:   time.Second * time.Duration(config.ConfigHub.Websocket.ConnWriteTimeout),
		MaxHeaderBytes: config.ConfigHub.Websocket.MaxHeaderLen,
	}
	httpServer.ListenAndServe()
}

type WsHandler struct {
	sv       *Server
	upgrader websocket.Upgrader
}

func (h *WsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.ZapLog.Error("upgrade error", zap.Error(err))
		return
	}

	//3.2 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
	if h.sv.ConnMgr.Len() >= 100 {
		conn.Close()
	}

	//3.3 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
	dealConn := NewWsConnection(h.sv, conn, h.sv.GenID(), h.sv.msgHandler)

	//3.4 启动当前链接的处理业务
	go dealConn.Start()

}

func (s *Server) httpServe() {

}

//Stop 停止服务
func (s *Server) Stop() {
	fmt.Println("[STOP] server , name ", s.Name)

	//将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
	s.ConnMgr.ClearConn()
}

//Serve 运行服务
func (s *Server) Serve() {
	s.StartTcp() // 默认开启
	if config.ConfigHub.Websocket.Enable {
		s.StartWebsocket()
	}

	// TODO quic、kcp的支持

	//TODO Server.Serve() 是否在启动服务的时候 还要处理其他的事情呢 可以在这里添加

	//阻塞,否则主Go退出， listenner的go将会退出
	select {}
}

//AddRouter 路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
func (s *Server) AddRouter(msgID uint32, router anet.Router) {
	s.msgHandler.AddRouter(msgID, router)
}


//AddRouter 路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
func (s *Server) SetRpcxRouter(router anet.RpcxRouter) {
	s.msgHandler.SetRpcxRouter(router)
}

//GetConnMgr 得到链接管理
func (s *Server) GetConnMgr() anet.ConnManager {
	return s.ConnMgr
}

//SetOnConnStart 设置该Server的连接创建时Hook函数
func (s *Server) SetOnConnStart(hookFunc func(anet.Connection)) {
	s.OnConnStart = hookFunc
}

//SetOnConnStop 设置该Server的连接断开时的Hook函数
func (s *Server) SetOnConnStop(hookFunc func(anet.Connection)) {
	s.OnConnStop = hookFunc
}

//CallOnConnStart 调用连接OnConnStart Hook函数
func (s *Server) CallOnConnStart(conn anet.Connection) {
	if s.OnConnStart != nil {
		fmt.Println("---> CallOnConnStart....")
		s.OnConnStart(conn)
	}
}

//CallOnConnStop 调用连接OnConnStop Hook函数
func (s *Server) CallOnConnStop(conn anet.Connection) {
	if s.OnConnStop != nil {
		fmt.Println("---> CallOnConnStop....")
		s.OnConnStop(conn)
	}
}

func (s *Server) Packet() anet.Packet {
	return s.packet
}

func printLogo() {
	//fmt.Println(zinxLogo)
	//fmt.Println(topLine)
	//fmt.Println(fmt.Sprintf("%s [Github] https://github.com/aceld                 %s", borderLine, borderLine))
	//fmt.Println(fmt.Sprintf("%s [tutorial] https://www.kancloud.cn/aceld/zinx     %s", borderLine, borderLine))
	//fmt.Println(bottomLine)
	//fmt.Printf("[Zinx] Version: %s, MaxConn: %d, MaxPacketSize: %d\n",
	//	utils.GlobalObject.Version,
	//	utils.GlobalObject.MaxConn,
	//	utils.GlobalObject.MaxPacketSize)
}

func init() {
}
