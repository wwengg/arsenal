// @Title
// @Description
// @Author  Wangwengang  2021/8/19 下午12:21
// @Update  Wangwengang  2021/8/19 下午12:21
package conf

type Rpcx struct {
	BasePath string `mapstructure:"base-path" json:"basePath" yaml:"base-path"` // basePath
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`               // ip
	Network  string `mapstructure:"network" json:"network" yaml:"network"`      // network
	Register string `mapstructure:"register" json:"register" yaml:"register"`   // 注册中心

}
