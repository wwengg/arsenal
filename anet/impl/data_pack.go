// @Title
// @Description
// @Author  Wangwengang  2021/8/17 下午5:25
// @Update  Wangwengang  2021/8/17 下午5:25
package impl

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/wwengg/arsenal/anet"
	"github.com/wwengg/arsenal/config"
)

var defaultHeaderLen uint32 = 16

//DataPack 封包拆包类实例，暂时不需要成员
type DataPack struct{}

//NewDataPack 封包拆包实例初始化方法
func NewDataPack() anet.Packet {
	return &DataPack{}
}

//GetHeadLen 获取包头长度方法
func (dp *DataPack) GetHeadLen() uint32 {
	//DataLen uint32(4字节) +  Op uint32(4字节) +  ver uint32(4字节) +  seq uint32(4字节)
	return defaultHeaderLen
}

//Pack 封包方法(压缩数据)
func (dp *DataPack) Pack(msg anet.Message) ([]byte, error) {
	//创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	//写dataLen
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	//写msgID
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}

	// 写版本ID
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetVer()); err != nil {
		return nil, err
	}

	// 写包ID
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetSeq()); err != nil {
		return nil, err
	}

	//写data数据
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

//Unpack 拆包方法(解压数据)
func (dp *DataPack) Unpack(binaryData []byte) (anet.Message, error) {
	//创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)

	//只解压head的信息，得到dataLen和msgID
	msg := &Message{}

	//读dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	//读op
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Op); err != nil {
		return nil, err
	}

	//读Ver
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Ver); err != nil {
		return nil, err
	}

	//读Seq
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Seq); err != nil {
		return nil, err
	}


	//判断dataLen的长度是否超出我们允许的最大包长度
	if config.ConfigHub.TcpConfig.MaxPacketSize > 0 && msg.DataLen > config.ConfigHub.TcpConfig.MaxPacketSize {
		return nil, errors.New("too large msg data received")
	}

	//这里只需要把head的数据拆包出来就可以了，然后再通过head的长度，再从conn读取一次数据
	return msg, nil
}
