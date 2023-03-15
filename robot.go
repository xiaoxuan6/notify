package notify

import (
    dinging_talk "github.com/xiaoxuan6/ding-talk"
    "github.com/xiaoxuan6/notify/v3/feishu"
    "github.com/xiaoxuan6/notify/v3/phprm"
    "github.com/xiaoxuan6/notify/v3/push_plus"
    "github.com/xiaoxuan6/notify/v3/server"
    "github.com/xiaoxuan6/notify/v3/utils"
    "github.com/xiaoxuan6/notify/v3/wechat"
    "github.com/xiaoxuan6/notify/v3/zhixi"
)

func NewDingTalk(config *utils.Config) *dinging_talk.Robot {
    return NewNotify(config).DingDing
}

func NewFeiShu(config *utils.Config) *feishu.Robot {
    return NewNotify(config).Feishu
}

func NewPhprm(config *utils.Config) *phprm.Robot {
    return NewNotify(config).Phprm
}

func NewPushPlus(config *utils.Config) *push_plus.Robot {
    return NewNotify(config).PushPlus
}

func NewServer(config *utils.Config) *server.Robot {
    return NewNotify(config).Server
}

func NewWechat(config *utils.Config) *wechat.Robot {
    return NewNotify(config).Wechat
}

func NewZhiXi(config *utils.Config) *zhixi.Robot {
    return NewNotify(config).ZhiXi
}
