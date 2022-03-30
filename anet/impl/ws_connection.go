// @Title
// @Description
// @Author  Wangwengang  2021/8/19 下午11:27
// @Update  Wangwengang  2021/8/19 下午11:27
package impl

import (
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/gorilla/websocket"

	"github.com/wwengg/arsenal/anet"
	"github.com/wwengg/arsenal/config"
)

type WsConnection struct {
	BaseConnection
	//当前连接的socket套接字
	Conn *websocket.Conn
}

//NewConnection 创建连接的方法
func NewWsConnection(server anet.Server, conn *websocket.Conn, connID uint64, msgHandler anet.MsgHandle) anet.Connection {
	//初始化Conn属性
	c := &WsConnection{
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
func (c *WsConnection) Write(data []byte) error {
	if err := c.Conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
		fmt.Println("Send Websocket Data error:, ", err)
		return err
	}
	return nil
}

//GetReader 读消息Goroutine，用于从客户端中读取数据
func (c *WsConnection) GetReader() (r io.Reader, err error) {
	messageType, r, err := c.Conn.NextReader()
	if err != nil {
		fmt.Println("websocket read msg error ", err)
		return nil, err
	}
	if websocket.BinaryMessage != messageType {
		fmt.Println("messageType != websocket.BinaryMessage")
		return nil, errors.New("messageType != websocket.BinaryMessage")
	}
	return r, nil
}

func (c *WsConnection) ConnClose() {
	c.Conn.Close()
}

//RemoteAddr 获取远程客户端地址信息
func (c *WsConnection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
