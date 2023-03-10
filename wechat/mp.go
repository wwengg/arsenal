// @Title  
// @Description  
// @Author  Wangwengang  2023/2/28 下午3:15
// @Update  Wangwengang  2023/2/28 下午3:15
package wechat

import (
	"github.com/chanxuehong/wechat/mp/core"

	"github.com/wwengg/arsenal/config"
)

var (
	AccessTokenServer core.AccessTokenServer = core.NewDefaultAccessTokenServer(config.ConfigHub.Wechat.AppId, config.ConfigHub.Wechat.WxAppSecret, nil)
	MP_CLIENT *core.Client = core.NewClient(AccessTokenServer,nil)
)
