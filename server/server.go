package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	talk "github.com/xiaoxuan6/ding-talk"
	"github.com/xiaoxuan6/notify/v3/feishu"
	"github.com/xiaoxuan6/notify/v3/wechat"
	"strings"
)

type Robot struct {
	Webhook string `json:"webhook"`
	Secret  string `json:"secret"`
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

func (r *Robot) SetSecret(secret string) *Robot {
	r.Secret = secret
	return r
}

func (r *Robot) Send(title, desp string) (item map[string]interface{}, err error) {

	if len(r.Webhook) < 1 {
		return item, errors.New("webhook 不能为空")
	}

	switch r.Channel {
	case "1":
		item, err = r.sendWechatTalk(desp)
	case "2":
		item, err = r.sendDingTalk(desp)
	case "3":
		item, err = r.sendFeishuTalk(desp)
	case "9":
		item, err = r.sendServe(title, desp, r.Channel)
	default:
		item, err = r.sendServe(title, desp, r.Channel)
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

	robot := wechat.NewRobot(key)
	err = robot.SendText(desp, []string{}, []string{})

	return item, err
}

func (r *Robot) sendDingTalk(desp string) (item map[string]interface{}, err error) {
	index := strings.LastIndex(r.Webhook, "=")
	accessToken := r.Webhook[index+1:]

	robot := talk.NewRobot(accessToken)

	if len(r.Secret) > 0 {
		robot.SetSecret(r.Secret)
	}

	err = robot.SendText(desp, []string{}, []string{}, false)

	return item, err
}

func (r *Robot) sendFeishuTalk(desc string) (item map[string]interface{}, err error) {
	index := strings.LastIndex(r.Webhook, "/")
	key := r.Webhook[index+1:]

	robot := feishu.NewRobot(key)

	if len(r.Secret) > 0 {
		robot.SetSecret(r.Secret)
	}

	err = robot.SendText(desc)

	return item, err
}
