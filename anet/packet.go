// @Title  
// @Description  
// @Author  Wangwengang  2021/8/17 下午5:12
// @Update  Wangwengang  2021/8/17 下午5:12
package anet

type Packet interface {
	Unpack(binaryData []byte) (Message, error)
	Pack(msg Message) ([]byte, error)
	GetHeadLen() uint32
}
