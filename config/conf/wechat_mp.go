// @Title  
// @Description  
// @Author  Wangwengang  2023/2/28 下午3:05
// @Update  Wangwengang  2023/2/28 下午3:05
package conf

type WechatMp struct {
	AliasName    string `mapstructure:"alias-name" json:"alias-name" yaml:"alias-name"`
	AppId string `mapstructure:"app-id" json:"appId" yaml:"appId"`
	WxAppSecret     string `mapstructure:"wx-app-secret" json:"wxAppSecret" yaml:"wxAppSecret"`
	WxOriId  string `mapstructure:"wx-ori-id" json:"wxOriId" yaml:"wxOriId"`
	WxToken	string `mapstructure:"wx-token" json:"wxToken" yaml:"wxToken"`
	WxEncodedAESKey string `mapstructure:"wx-encoded-AESKey" json:"wxEncodedAESKey" yaml:"wxEncodedAESKey"`
}
