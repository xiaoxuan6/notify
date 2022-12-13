package wechat

type textMessage struct {
	MsgType string `json:"msgtype"`
	Text    text   `json:"text"`
}

type text struct {
	Content             string   `json:"content" describe:"文本内容，最长不超过2048个字节，必须是utf8编码"`
	MentionedList       []string `json:"mentioned_list" describe:"userid的列表，提醒群中的指定成员(@某个成员)，@all表示提醒所有人"`
	MentionedMobileList []string `json:"mentioned_mobile_list" describe:"手机号列表，提醒手机号对应的群成员(@某个成员)，@all表示提醒所有人"`
}

type markdownMessage struct {
	MsgType  string   `json:"msgtype"`
	Markdown markdown `json:"markdown"`
}

type markdown struct {
	Content string `json:"content" describe:"markdown内容，最长不超过4096个字节，必须是utf8编码"`
}

type imageMessage struct {
	MsgType string `json:"msgtype"`
	Image   image  `json:"image"`
}

type image struct {
	Base64 string `json:"base64" describe:"图片内容的base64编码"`
	Md5    string `json:"md5" describe:"图片内容（base64编码前）的md5值"`
}

type newsMessage struct {
	MsgType string `json:"msgtype"`
	News    news   `json:"news"`
}

type news struct {
	Articles []Articles `json:"articles" describe:"图文消息，一个图文消息支持1到8条图文"`
}

type Articles struct {
	Title       string `json:"title" describe:"标题，不超过128个字节，超过会自动截断"`
	Description string `json:"description" describe:"描述，不超过512个字节，超过会自动截断"`
	Url         string `json:"url" describe:"点击后跳转的链接。"`
	PicUrl      string `json:"picurl" describe:"图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图 1068*455，小图150*150。"`
}

type fileMessage struct {
	MsgType string `json:"msgtype"`
	File    file   `json:"file" describe:"文件id，通过下文的文件上传接口获取"`
}

type file struct {
	MediaId string `json:"media_id" describe:"文件id，通过下文的文件上传接口获取"`
}
