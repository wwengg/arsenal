// @Title  
// @Description  
// @Author  Wangwengang  2021/8/19 上午1:50
// @Update  Wangwengang  2021/8/19 上午1:50
package conf

type TcpConfig struct {
	Name			string	`mapstructure:"name" json:"name" yaml:"name"`                         // serverName
	Ip           	string `mapstructure:"ip" json:"ip" yaml:"ip"`                                // ip
	TcpPort         int    `mapstructure:"port" json:"port" yaml:"port"`                          // 端口值
	MaxConn			int	   `mapstructure:"max-conn" json:"maxConn" yaml:"max-conn"`               // 最大连接数
	WorkerPoolSize	uint32	   `mapstructure:"worker-pool-size" json:"workerPoolSize" yaml:"worker-pool-size"`    // 工作池数量
	MaxWorkerTaskLen uint32	`mapstructure:"max-worker-task-len" json:"maxWorkerTaskLen" yaml:"max-worker-task-len"`
	MaxMsgChanLen	uint32		`mapstructure:"max-msg-chan-len" json:"maxMsgChanLen" yaml:"max-msg-chan-len"`
	MaxPacketSize	uint32		`mapstructure:"max-packet-size" json:"maxPacketSize" yaml:"max-packet-size"`
}
