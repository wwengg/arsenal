// @Title  
// @Description  
// @Author  Wangwengang  2021/8/17 下午5:11
// @Update  Wangwengang  2021/8/17 下午5:11
package internal

import (
	"github.com/wwengg/arsenal/anet"
)

type Option func(s *TcpServer)

// 只要实现Packet 接口可自由实现数据包解析格式，如果没有则使用默认解析格式
func WithPacket(pack anet.Packet) Option {
	return func(s *TcpServer) {
		s.packet = pack
	}
}

