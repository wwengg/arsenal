// @Title
// @Description
// @Author  Wangwengang  2021/8/17 下午5:07
// @Update  Wangwengang  2021/8/17 下午5:07
package anet

type Server interface {
	StartTcp()                             // 启动服务器方法
	StartWebsocket()                       // StartWebsocket
	Stop()                                 // 停止服务器方法
	Serve()                                //开启业务服务方法
	AddRouter(msgID uint32, router Router) //路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
	SetRpcxRouter(router RpcxRouter)
	GetConnMgr() ConnManager               //得到链接管理
	SetOnConnStart(func(Connection))       //设置该Server的连接创建时Hook函数
	SetOnConnStop(func(Connection))        //设置该Server的连接断开时的Hook函数
	CallOnConnStart(conn Connection)       //调用连接OnConnStart Hook函数
	CallOnConnStop(conn Connection)        //调用连接OnConnStop Hook函数
	Packet() Packet
}
