package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	talk "github.com/xiaoxuan6/ding-talk"
	"github.com/xiaoxuan6/notify/feishu"
	wechat_talk "github.com/xiaoxuan6/wechat-talk"
	"net/url"
	"strings"
)

type Robot struct {
	Webhook string `json:"webhook"`
	Channel string `json:"channel"`
}

func NewRobot(webhook string) *Robot {
	return &Robot{
		Webhook: webhook,
	}
}

func (r *Robot) SetChannel(channel string) *Robot {
	r.Channel = channel
	return r
}

func (r *Robot) Send(title, desp string) (item map[string]interface{}, err error) {

	channel := r.Channel
	if channel == "" {
		channel = "9" // 使用默认通道
	}

	switch channel {
	case "1":
		item, err = r.sendWechatTalk(desp)
	case "2":
		item, err = r.sendDingTalk(desp)
	case "3":
		item, err = r.sendFeishuTalk(desp)
	case "9":
		item, err = r.sendServe(title, desp, channel)
	}

	return
}

func (r *Robot) sendServe(title, desp, channel string) (map[string]interface{}, error) {
	formData := map[string]string{
		"title":   title,
		"desp":    desp,
		"channel": channel,
	}
	resp, err := resty.New().R().SetFormData(formData).Post(r.Webhook)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("请求失败：%s", err.Error()))
	}

	var response map[string]interface{}
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil, errors.New("json 解析数据失败")
	}

	if response["code"] != float64(0) {
		return nil, errors.New(response["message"].(string))
	}

	return response["data"].(map[string]interface{}), nil
}

func (r *Robot) sendWechatTalk(desp string) (item map[string]interface{}, err error) {
	index := strings.LastIndex(r.Webhook, "=")
	key := r.Webhook[index+1:]

	robot := wechat_talk.NewRobot(key)
	err = robot.SendText(desp, []string{}, []string{})

	return item, err
}

func (r *Robot) sendDingTalk(desp string) (item map[string]interface{}, err error) {
	index := strings.LastIndex(r.Webhook, "?")
	urls := r.Webhook[index+1:]
	val, _ := url.ParseQuery(urls)

	robot := talk.NewRobot(val.Get("access_token"))
	if server := val.Get("server"); server != "" {
		robot.SetSecret(server)
	}
	err = robot.SendText(desp, []string{}, []string{}, false)

	return item, err
}

func (r *Robot) sendFeishuTalk(desc string) (item map[string]interface{}, err error) {
	index := strings.LastIndex(r.Webhook, "/")
	endIndex := strings.LastIndex(r.Webhook, "?")

	var key string
	if endIndex > 0 {
		key = r.Webhook[index+1 : endIndex]
	} else {
		key = r.Webhook[index+1:]
	}

	robot := feishu.NewRobot(key)
	val, _ := url.ParseQuery(r.Webhook[endIndex+1:])
	if server := val.Get("server"); server != "" {
		robot.SetSecret(server)
	}
	err = robot.SendText(desc)

	return item, err
}
