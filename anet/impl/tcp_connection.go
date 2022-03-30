// @Title
// @Description
// @Author  Wangwengang  2021/8/17 下午4:45
// @Update  Wangwengang  2021/8/17 下午4:45
package impl

import (
	"fmt"
	"io"
	"net"

	"github.com/wwengg/arsenal/anet"
	"github.com/wwengg/arsenal/config"
)

type TcpConnection struct {
	BaseConnection
	//当前连接的socket TCP套接字
	Conn *net.TCPConn
}

//NewConnection 创建连接的方法
func NewConnection(server anet.Server, conn *net.TCPConn, connID uint64, msgHandler anet.MsgHandle) anet.Connection {
	//初始化Conn属性
	c := &TcpConnection{
		BaseConnection: BaseConnection{
			TCPServer:   server,
			ConnID:      connID,
			isClosed:    false,
			MsgHandler:  msgHandler,
			msgChan:     make(chan []byte),
			msgBuffChan: make(chan []byte, config.ConfigHub.TcpConfig.MaxMsgChanLen),
			property:    nil,
		},
		Conn: conn,
	}

	//将新创建的Conn添加到链接管理中
	c.TCPServer.GetConnMgr().Add(c)
	return c
}

//Write 写消息Goroutine， 用户将数据发送给客户端
func (c *TcpConnection) Write(data []byte) error {
	if _, err := c.Conn.Write(data); err != nil {
		fmt.Println("Send TCP Data error:, ", err)
		return err
	}
	return nil
}

//GetReader 读消息Goroutine，用于从客户端中读取数据
func (c *TcpConnection) GetReader() (r io.Reader, err error) {
	return c.Conn, nil
}

func (c *TcpConnection) ConnClose() {
	c.Conn.Close()
}

//RemoteAddr 获取远程客户端地址信息
func (c *TcpConnection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
