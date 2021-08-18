// @Title  
// @Description  
// @Author  Wangwengang  2021/8/17 下午5:23
// @Update  Wangwengang  2021/8/17 下午5:23
package anet

/*
	封包数据和拆包数据
	直接面向TCP连接中的数据流,为传输数据添加头部信息，用于处理TCP粘包问题。
*/
type DataPack interface {
	GetHeadLen() uint32                //获取包头长度方法
	Pack(msg Message) ([]byte, error) //封包方法
	Unpack([]byte) (Message, error)   //拆包方法
}

