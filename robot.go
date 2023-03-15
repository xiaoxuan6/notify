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
    return newNotify(config).DingDing
}

func NewFeiShu(config *utils.Config) *feishu.Robot {
    return newNotify(config).Feishu
}

func NewPhprm(config *utils.Config) *phprm.Robot {
    return newNotify(config).Phprm
}

func NewPushPlus(config *utils.Config) *push_plus.Robot {
    return newNotify(config).PushPlus
}

func NewServer(config *utils.Config) *server.Robot {
    return newNotify(config).Server
}

func NewWechat(config *utils.Config) *wechat.Robot {
    return newNotify(config).Wechat
}

func NewZhiXi(config *utils.Config) *zhixi.Robot {
    return newNotify(config).ZhiXi
}
