// @Title
// @Description
// @Author  Wangwengang  2022/2/17 上午11:19
// @Update  Wangwengang  2022/2/17 上午11:19
package conf

type HttpGateway struct {
	Env  string `mapstructure:"env" json:"env" yaml:"env"`
	Addr int    `mapstructure:"addr" json:"addr" yaml:"addr"`
}
