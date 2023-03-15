package utils

type Config struct {
	Dinging  DingingConfig  `json:"dinging" yaml:"dinging"`
	Wechat   WechatConfig   `json:"wechat" yaml:"wechat"`
	Feishu   FeishuConfig   `json:"feishu" yaml:"feishu"`
	Server   ServerConfig   `json:"server" yaml:"server"`
	PushPlus PushPlusConfig `json:"push_plus" yaml:"push_plus"`
	Phprm    Phprm          `json:"phprm" yaml:"phprm"`
}

type DingingConfig struct {
	AccessToken string `json:"access_token" yaml:"access_token"`
	Secret      string `json:"secret" yaml:"secret"`
}

type WechatConfig struct {
	Key string `json:"key" yaml:"key"`
}

type FeishuConfig struct {
	AccessToken string `json:"access_token" yaml:"access_token"`
	Secret      string `json:"secret" yaml:"secret"`
}

const (
	WechatChannel   = "1" // 企业微信群机器人
	DingingChannel  = "2" // 钉钉群机器人
	FeishuChannel   = "3" // 飞书群机器人
	FangtangChannel = "9" // 方糖服务号:默认使用
)

type ServerConfig struct {
	Webhook string `json:"webhook" yaml:"webhook"`
	Secret  string `json:"secret" yaml:"secret"`
	Channel string `json:"channel" yaml:"channel"`
}

type PushPlusConfig struct {
	Token string `json:"token" yaml:"token"`
}

type Phprm struct {
	Token string `json:"token" yaml:"token"`
}
