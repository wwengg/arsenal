// @Title  
// @Description  
// @Author  Wangwengang  2021/8/17 下午9:54
// @Update  Wangwengang  2021/8/17 下午9:54
package internal

import (
	"github.com/wwengg/arsenal/anet"
)

//Request 请求
type Request struct {
	conn anet.Connection //已经和客户端建立好的 链接
	msg  anet.Message    //客户端请求的数据
}

//GetConnection 获取请求连接信息
func (r *Request) GetConnection() anet.Connection {
	return r.conn
}

//GetData 获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

//GetMsgID 获取请求的消息的ID
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}

