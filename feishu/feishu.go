package feishu

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/xiaoxuan6/notify/v2/utils"
	"time"
)

const webhook = "https://open.feishu.cn/open-apis/bot/v2/hook/"

const (
	MsgTypeText        = "text"
	MsgTypePost        = "post"       // 发送富文本消息
	MsgTypeShareChat   = "share_chat" // 发送群名片
	MsgTypeImage       = "image"
	MsgTypeInteractive = "interactive" // 发送消息卡片
)

type Robot struct {
	accessToken string
	secret      string
}

func NewRobot(accessToken string) *Robot {
	return &Robot{
		accessToken: accessToken,
	}
}

func (r *Robot) SetSecret(secret string) *Robot {
	r.secret = secret
	return r
}

// SendText 发送文本消息
func (r *Robot) SendText(text string) error {
	return r.send(&textMessage{
		MsgType: MsgTypeText,
		Content: content{
			Text: text,
		},
	})
}

// SendPost 发送富文本消息
func (r *Robot) SendPost(title string, content [][]ContentParams) error {
	return r.send(&postMessage{
		MsgType: MsgTypePost,
		Content: contentPost{
			Post: post{
				ZhCn: zhCn{
					Title:   title,
					Content: content,
				},
			},
		},
	})
}

//SendShareChat 发送群名片
func (r *Robot) SendShareChat(chatId string) error {
	return r.send(&shareChatMessage{
		MsgType: MsgTypeShareChat,
		Content: sContent{
			ShareChatId: chatId,
		},
	})
}

// SendImage 发送图片
func (r *Robot) SendImage(imageKey string) error {
	return r.send(&imageMessage{
		MsgType: MsgTypeImage,
		Content: iContent{
			ImageKey: imageKey,
		},
	})
}

//func (r *Robot) SendInteractive() error {
//	return r.send()
//}

func (r *Robot) send(msg interface{}) error {

	if len(r.secret) > 0 {
		msg = genSigned(r.secret, msg)
	}

	if len(r.accessToken) < 1 {
		return errors.New("access_token 不能为空")
	}

	marshal, err := json.Marshal(msg)
	if err != nil {
		return errors.New("json 格式化数据失败")
	}

	url := fmt.Sprintf("%s%s", webhook, r.accessToken)
	resp, err := resty.New().R().SetBody(string(marshal)).Post(url)
	if err != nil {
		return errors.New(fmt.Sprintf("请求失败：%s", err.Error()))
	}

	var item map[string]interface{}
	err = json.Unmarshal(resp.Body(), &item)
	if err != nil {
		return errors.New("json 解析数据失败")
	}

	if _, ok := item["code"]; ok {
		return errors.New(item["msg"].(string))
	}

	if item["StatusCode"] != float64(0) {
		return errors.New(item["StatusMessage"].(string))
	}

	return nil
}

func genSigned(secret string, msg interface{}) (data map[string]interface{}) {
	timestamp := time.Now().Unix()
	sign := utils.GenSignedByHmacSHA256(secret, timestamp)

	b, _ := json.Marshal(msg)
	_ = json.Unmarshal(b, &data)

	data["timestamp"] = timestamp
	data["sign"] = sign

	return
}
