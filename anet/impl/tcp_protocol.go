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
)

type TcpProtocol struct {
	//当前连接的socket TCP套接
	Conn *net.TCPConn
}

// NewTcpProtocol
func NewTcpProtocol (conn *net.TCPConn) anet.Protocol{
	return &TcpProtocol{Conn: conn}
}

//Write 写消息Goroutine， 用户将数据发送给客户端
func (c *TcpProtocol) Write(data []byte) error {
	if _, err := c.Conn.Write(data); err != nil {
		fmt.Println("Send TCP Data error:, ", err)
		return err
	}
	return nil
}

//GetReader 读消息Goroutine，用于从客户端中读取数据
func (c *TcpProtocol) GetReader() (r io.Reader, err error) {
	return c.Conn, nil
}

func (c *TcpProtocol) ConnClose() {
	c.Conn.Close()
}

//RemoteAddr 获取远程客户端地址信息
func (c *TcpProtocol) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
