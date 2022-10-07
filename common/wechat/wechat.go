package wechat

import (
	"errors"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram/auth"
	mpConfig "github.com/silenceper/wechat/v2/miniprogram/config"

	"github.com/njupt-sast/atsast-apply-module-server/config"
)

var (
	wc          = wechat.NewWechat()
	miniProgram = wc.GetMiniProgram(&mpConfig.Config{
		AppID:     config.Wechat.AppId,
		AppSecret: config.Wechat.AppSecret,
		Cache:     cache.NewMemory(),
	})
)

func Code2Session(jsCode *string) (*auth.ResCode2Session, error) {
	result, err := miniProgram.GetAuth().Code2Session(*jsCode)
	if err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, nil
}
