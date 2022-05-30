package notify

import (
	dinging_talk "github.com/xiaoxuan6/ding-talk"
	"github.com/xiaoxuan6/notify/v2/dinging"
	"github.com/xiaoxuan6/notify/v2/feishu"
	"github.com/xiaoxuan6/notify/v2/server"
	"github.com/xiaoxuan6/notify/v2/utils"
	"github.com/xiaoxuan6/notify/v2/wechat"
	wechat_talk "github.com/xiaoxuan6/wechat-talk"
)

type Notify struct {
	DingDing *dinging_talk.Robot
	Wechat   *wechat_talk.Robot
	Feishu   *feishu.Robot
	Server   *server.Robot
}

func NewNotify(config *utils.Config) *Notify {

	notify := &Notify{}

	notify.DingDing = dinging.RegisterProvider(config)

	notify.Wechat = wechat.RegisterProvider(config)

	notify.Feishu = feishu.RegisterProvider(config)

	notify.Server = server.RegisterProvider(config)

	return notify
}
