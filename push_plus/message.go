package push_plus

const (
	TEMPLATE_HTML     = "html"     // 支持html文本。为空默认使用html模板
	TEMPLATE_TXT      = "txt"      // 纯文本内容,不转义html内容,换行使用\n
	TEMPLATE_JSON     = "json"     // 可视化展示json格式内容
	TEMPLATE_MARKDOWN = "markdown" // 内容基于markdown格式展示
)

const (
	CHANNEL_WECHAT  = "wechat"  // 微信公众号,默认发送渠道
	CHANNEL_WEBHOOK = "webhook" // 第三方webhook服务；企业微信机器人、钉钉机器人、飞书机器人
	CHANNEL_CP      = "cp"      // 企业微信应用
	CHANNEL_MAIL    = "mail"    // 邮件
	CHANNEL_SMS     = "sms"     // 短信；收费使用，1条短信扣减10积分
)

type Message struct {
	Token       string `json:"token"`
	Title       string `json:"title" describe:"消息标题"`
	Content     string `json:"content" describe:"具体消息内容，根据不同template支持不同格式"`
	Template    string `json:"template" describe:"发送消息模板"`
	Channel     string `json:"channel" describe:"发送渠道"`
	Webhook     string `json:"webhook" describe:"webhook编码"`
	CallbackUrl string `json:"callback_url" describe:"回调地址，异步回调发送结果"`
	TimeStamp   string `json:"time_stamp" describe:"时间戳，毫秒。如小于当前时间，消息将无法发送"`
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}
