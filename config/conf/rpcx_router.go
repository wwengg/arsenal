// @Title  
// @Description  
// @Author  Wangwengang  2021/8/26 上午10:09
// @Update  Wangwengang  2021/8/26 上午10:09
package conf

type RpcxRouter struct {
	Op uint32	`mapstructure:"op" json:"op" yaml:"op"`
	ServicePath  string		`mapstructure:"service-path" json:"servicePath" yaml:"service-path"`
	ServiceMethod  string		`mapstructure:"service-method" json:"serviceMethod" yaml:"service-method"`
	Oneway  bool		`mapstructure:"oneway" json:"oneway" yaml:"oneway"`
	BackOp	uint32	`mapstructure:"back-op" json:"backOp" yaml:"back-op"`
}
