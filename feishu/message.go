package feishu

type textMessage struct {
	MsgType string  `json:"msg_type"`
	Content content `json:"content"`
}

type content struct {
	Text string `json:"text"`
}

type postMessage struct {
	MsgType string      `json:"msg_type"`
	Content contentPost `json:"content"`
}

type contentPost struct {
	Post post `json:"post"`
}

type post struct {
	ZhCn zhCn `json:"zh_cn"`
}

type zhCn struct {
	Title   string            `json:"title"`
	Content [][]ContentParams `json:"content"`
}

type ContentParams struct {
	Tag    string `json:"tag"`
	Text   string `json:"text"`
	Href   string `json:"href"`
	UserId string `json:"user_id"`
}

type shareChatMessage struct {
	MsgType string   `json:"msg_type"`
	Content sContent `json:"content"`
}

type sContent struct {
	ShareChatId string `json:"share_chat_id"`
}

type imageMessage struct {
	MsgType string   `json:"msg_type"`
	Content iContent `json:"content"`
}

type iContent struct {
	ImageKey string `json:"image_key"`
}

/*
type interactiveMessage struct {
	MsgType string `json:"msg_type"`
	Card    card   `json:"card"`
}

type card struct {
	Config   config     `json:"config"`
	Elements []Elements `json:"elements"`
	Header   header     `json:"header" describe:"在header中描述卡片的标题"`
}

type config struct {
	WideScreenMode bool `json:"wide_screen_mode"`
	EnableForward  bool `json:"enable_forward"`
}

type Elements struct {
	Tag     string    `json:"tag"`
	Text    Text      `json:"text"`
	Actions []Actions `json:"actions"`
}

type Text struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type Actions struct {
	Tag   string   `json:"tag"`
	Text  Text     `json:"text"`
	Url   string   `json:"url"`
	Type  string   `json:"type"`
	Value struct{} `json:"value"`
}

type header struct {
	Title    title  `json:"title" describe:"配置卡片标题内容"`
	Template string `json:"template" describe:"卡片标题的主题色"`
}

type title struct {
	Content string `json:"content" describe:"卡片标题文案内容"`
	Tag     string `json:"tag" describe:"仅支持plain_text"`
}
*/
