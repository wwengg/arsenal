// @Title
// @Description
// @Author  Wangwengang  2021/8/20 下午2:38
// @Update  Wangwengang  2021/8/20 下午2:38
package conf

type Websocket struct {
	Enable           bool   `mapstructure:"enable" json:"enable" yaml:"enable"` // 是否开启Websocket
	Addr             string `mapstructure:"addr" json:"addr" yaml:"addr"`
	ConnWriteTimeout int    `mapstructure:"conn-write-timeout" json:"connWriteTimeout" yaml:"conn-write-timeout"`
	ConnReadTimeout  int    `mapstructure:"conn-read-timeout" json:"connReadTimeout" yaml:"conn-read-timeout"`
	MaxHeaderLen     int    `mapstructure:"max-header-len" json:"maxHeaderLen" yaml:"max-header-len"`
}
