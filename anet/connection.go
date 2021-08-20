// @Title  Connection interface
// @Description  tcp websocket
// @Author  Wangwengang  2021/8/17 下午4:22
// @Update  Wangwengang  2021/8/17 下午4:22
package anet

import (
	"context"
	"net"

	"github.com/gorilla/websocket"
)

type Connection interface {
	Start()                   //启动连接，让当前连接开始工作
	Stop()                    //停止连接，结束当前连接状态M
	Context() context.Context //返回ctx，用于用户自定义的go程获取连接退出状态

	GetTcpConnection() *net.TCPConn   //从当前连接获取原始的socket TCPConn
	GetWsConnection() *websocket.Conn // 从当前连接获取原始的websocket Conn
	GetConnID() uint64                //获取当前连接ID
	RemoteAddr() net.Addr             //获取远程客户端地址信息

	SendMsg(msgID uint32, data []byte) error     //直接将Message数据发送数据给远程的TCP客户端(无缓冲)
	SendBuffMsg(msgID uint32, data []byte) error //直接将Message数据发送给远程的TCP客户端(有缓冲)

	SetProperty(key string, value interface{})   //设置链接属性
	GetProperty(key string) (interface{}, error) //获取链接属性
	RemoveProperty(key string)                   //移除链接属性
}
