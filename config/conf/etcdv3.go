// @Title
// @Description
// @Author  Wangwengang  2021/8/19 下午12:42
// @Update  Wangwengang  2021/8/19 下午12:42
package conf

type EtcdV3 struct {
	Addr []string `mapstructure:"addr" json:"addr" yaml:"addr"`
}
