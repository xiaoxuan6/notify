package server

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/go-resty/resty/v2"
    talk "github.com/xiaoxuan6/ding-talk"
    "github.com/xiaoxuan6/notify/v3/feishu"
    "github.com/xiaoxuan6/notify/v3/utils"
    "github.com/xiaoxuan6/notify/v3/wechat"
    "strings"
)

type Robot struct {
}


func (r *Robot) Send(title, desp string) (item map[string]interface{}, err error) {

    if len(utils.GlobalConfig.Server.Webhook) < 1 {
        return item, errors.New("webhook 不能为空")
    }

    channel := utils.GlobalConfig.Server.Channel
    switch channel {
    case "1":
        item, err = r.sendWechatTalk(desp)
    case "2":
        item, err = r.sendDingTalk(desp)
    case "3":
        item, err = r.sendFeishuTalk(desp)
    case "9":
        item, err = r.sendServe(title, desp, channel)
    default:
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
    resp, err := resty.New().R().SetFormData(formData).Post(utils.GlobalConfig.Server.Webhook)

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
    index := strings.LastIndex(utils.GlobalConfig.Server.Webhook, "=")
    key := utils.GlobalConfig.Server.Webhook[index+1:]

    utils.GlobalConfig.Wechat.Key = key

    robot := &wechat.Robot{}
    err = robot.SendText(desp, []string{}, []string{})

    return item, err
}

func (r *Robot) sendDingTalk(desp string) (item map[string]interface{}, err error) {
    index := strings.LastIndex(utils.GlobalConfig.Server.Webhook, "=")
    accessToken := utils.GlobalConfig.Server.Webhook[index+1:]

    robot := talk.NewRobot(accessToken)

    if len(utils.GlobalConfig.Server.Secret) > 0 {
        robot.SetSecret(utils.GlobalConfig.Server.Secret)
    }

    err = robot.SendText(desp, []string{}, []string{}, false)

    return item, err
}

func (r *Robot) sendFeishuTalk(desc string) (item map[string]interface{}, err error) {
    index := strings.LastIndex(utils.GlobalConfig.Server.Webhook, "/")
    key := utils.GlobalConfig.Server.Webhook[index+1:]

    utils.GlobalConfig.Feishu.AccessToken = key
    utils.GlobalConfig.Feishu.Secret = utils.GlobalConfig.Server.Secret

    robot := &feishu.Robot{}
    err = robot.SendText(desc)

    return item, err
}
