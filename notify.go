package notify

import (
	dinging_talk "github.com/xiaoxuan6/ding-talk"
	"github.com/xiaoxuan6/notify/v2/dinging"
	"github.com/xiaoxuan6/notify/v2/feishu"
	"github.com/xiaoxuan6/notify/v2/phprm"
	"github.com/xiaoxuan6/notify/v2/push_plus"
	"github.com/xiaoxuan6/notify/v2/server"
	"github.com/xiaoxuan6/notify/v2/utils"
	"github.com/xiaoxuan6/notify/v2/wechat"
)

type Notify struct {
	DingDing *dinging_talk.Robot
	Wechat   *wechat.Robot
	Feishu   *feishu.Robot
	Server   *server.Robot
	PushPlus *push_plus.Root
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
