package utils

type Config struct {
	Dinging  dingingConfig  `json:"dinging"`
	Wechat   wechatConfig   `json:"wechat"`
	Feishu   feishuConfig   `json:"feishu"`
	Server   serverConfig   `json:"server"`
	PushPlus pushPlusConfig `json:"push_plus"`
}

type dingingConfig struct {
	AccessToken string `json:"access_token"`
	Secret      string `json:"secret"`
}

type wechatConfig struct {
	Key string `json:"key"`
}

type feishuConfig struct {
	AccessToken string `json:"access_token"`
	Secret      string `json:"secret"`
}

const (
	WechatChannel   = "1" // 企业微信群机器人
	DingingChannel  = "2" // 钉钉群机器人
	FeishuChannel   = "3" // 飞书群机器人
	FangtangChannel = "9" // 方糖服务号:默认使用
)

type serverConfig struct {
	Webhook string `json:"webhook"`
	Secret  string `json:"secret"`
	Channel string `json:"channel"`
}

type pushPlusConfig struct {
	Token string `json:"token"`
}
