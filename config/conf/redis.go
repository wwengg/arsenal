// @Title  
// @Description  
// @Author  Wangwengang  2021/8/24 下午9:37
// @Update  Wangwengang  2021/8/24 下午9:37
package conf

type Redis struct {
	Addr []string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Db   int      `mapstructure:"db" json:"db" yaml:"db"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	MasterName string `mapstructure:"master-name" json:"masterName" yaml:"master-name"`
}
