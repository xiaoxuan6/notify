package utils

type Config struct {
	Dinging  DingingConfig  `json:"dinging"`
	Wechat   WechatConfig   `json:"wechat"`
	Feishu   FeishuConfig   `json:"feishu"`
	Server   ServerConfig   `json:"server"`
	PushPlus PushPlusConfig `json:"push_plus"`
	Phprm    Phprm          `json:"phprm"`
}

type DingingConfig struct {
	AccessToken string `json:"access_token"`
	Secret      string `json:"secret"`
}

type WechatConfig struct {
	Key string `json:"key"`
}

type FeishuConfig struct {
	AccessToken string `json:"access_token"`
	Secret      string `json:"secret"`
}

const (
	WechatChannel   = "1" // 企业微信群机器人
	DingingChannel  = "2" // 钉钉群机器人
	FeishuChannel   = "3" // 飞书群机器人
	FangtangChannel = "9" // 方糖服务号:默认使用
)

type ServerConfig struct {
	Webhook string `json:"webhook"`
	Secret  string `json:"secret"`
	Channel string `json:"channel"`
}

type PushPlusConfig struct {
	Token string `json:"token"`
}

type Phprm struct {
	Token string `json:"token"`
}
