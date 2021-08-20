// @Title
// @Description
// @Author  Wangwengang  2021/8/17 下午5:26
// @Update  Wangwengang  2021/8/17 下午5:26
package impl

import (
	"github.com/wwengg/proto/pigeon"
)

//Message 消息
type Message struct {
	pigeon.Proto
	//DataLen uint32 //消息的长度
	//ID      uint32 //消息的ID
	//Data    []byte //消息的内容
}

//NewMsgPackage 创建一个Message消息包
func NewMsgPackage(ID uint32, data []byte) *Message {
	return &Message{
		pigeon.Proto{
			DataLen:              uint32(len(data)),
			Op:                   ID,
			Ver:                  1,
			Seq:                  0,
			Body:                 data,
		},
	}
}

//GetDataLen 获取消息数据段长度
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

//GetMsgID 获取消息ID
func (msg *Message) GetMsgID() uint32 {
	return msg.Op
}

//GetVer 获取版本
func (msg *Message) GetVer() uint32 {
	return msg.Ver
}

//GetSeq 获取版本
func (msg *Message) GetSeq() uint32 {
	return msg.Seq
}

//GetData 获取消息内容
func (msg *Message) GetData() []byte {
	return msg.Body
}

//SetDataLen 设置消息数据段长度
func (msg *Message) SetDataLen(len uint32) {
	msg.DataLen = len
}

//SetMsgID 设计消息ID
func (msg *Message) SetMsgID(msgID uint32) {
	msg.Op = msgID
}

//SetData 设计消息内容
func (msg *Message) SetData(data []byte) {
	msg.Body = data
}
