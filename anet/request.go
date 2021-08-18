// @Title  
// @Description  
// @Author  Wangwengang  2021/8/17 下午5:14
// @Update  Wangwengang  2021/8/17 下午5:14
package anet

/*
	Request 接口：
	实际上是把客户端请求的链接信息 和 请求的数据 包装到了 Request里
*/
type Request interface {
	GetConnection() Connection //获取请求连接信息
	GetData() []byte            //获取请求消息的数据
	GetMsgID() uint32           //获取请求的消息ID
}
