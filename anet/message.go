// @Title  
// @Description  
// @Author  Wangwengang  2021/8/17 下午5:12
// @Update  Wangwengang  2021/8/17 下午5:12
package anet

type Message interface {
	GetDataLen() uint32 //获取消息数据段长度
	GetMsgID() uint32   //获取消息ID
	GetData() []byte    //获取消息内容

	SetMsgID(uint32)   //设计消息ID
	SetData([]byte)    //设计消息内容
	SetDataLen(uint32) //设置消息数据段长度
}
