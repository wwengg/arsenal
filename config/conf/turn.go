// @Title  
// @Description  
// @Author  Wangwengang  2022/5/2 下午5:59
// @Update  Wangwengang  2022/5/2 下午5:59
package conf

type Turn struct {
	PublicIp string `mapstructure:"public-ip" json:"publicIp" yaml:"public-ip"`
	Port     int `mapstructure:"port" json:"port" yaml:"port"`
	Realm  string `mapstructure:"realm" json:"realm" yaml:"realm"`
}
