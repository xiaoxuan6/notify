package wechat

import (
	"github.com/xiaoxuan6/notify/v2/utils"
	wechat_talk "github.com/xiaoxuan6/wechat-talk"
)

func RegisterProvider(config *utils.Config) *wechat_talk.Robot {

	return wechat_talk.NewRobot(config.Wechat.Key)

}
