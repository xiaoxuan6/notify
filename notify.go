package notify

import (
	dinging_talk "github.com/xiaoxuan6/ding-talk"
	"github.com/xiaoxuan6/notify/v3/dinging"
	"github.com/xiaoxuan6/notify/v3/feishu"
	"github.com/xiaoxuan6/notify/v3/phprm"
	"github.com/xiaoxuan6/notify/v3/push_plus"
	"github.com/xiaoxuan6/notify/v3/server"
	"github.com/xiaoxuan6/notify/v3/utils"
	"github.com/xiaoxuan6/notify/v3/wechat"
)

type Notify struct {
	DingDing *dinging_talk.Robot
	Wechat   *wechat.Robot
	Feishu   *feishu.Robot
	Server   *server.Robot
	PushPlus *push_plus.Robot
	Phprm    *phprm.Robot
}

func NewNotify(config *utils.Config) *Notify {

	notify := &Notify{}

	notify.DingDing = dinging.RegisterProvider(config)
	notify.Wechat = wechat.RegisterProvider(config)
	notify.Feishu = feishu.RegisterProvider(config)
	notify.Server = server.RegisterProvider(config)
	notify.PushPlus = push_plus.RegisterProvider(config)
	notify.Phprm = phprm.RegisterProvider(config)

	return notify
}
