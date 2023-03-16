package phprm

import (
	"encoding/json"
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
    "github.com/xiaoxuan6/notify/v3/utils"
    "net/url"
)

const URI = "https://www.phprm.com/services/push/trigger"

type Robot struct {
}

type Response struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Data    struct {
        MessageIdList []string `json:"messageIdList"`
    } `json:"data"`
}

func (r *Robot) Send(head, body string) (err error, res *Response) {

	requestParams := map[string]interface{}{
		"token": utils.GlobalConfig.Phprm.Token,
		"head":  head,
		"body":  body,
	}
	if err = validation.Validate(requestParams, validation.Map(
		validation.Key("token", validation.Required.Error("invalid token")),
		validation.Key("head", validation.Required.Error("request head param required"), validation.NilOrNotEmpty.Error("request head param not empty")),
		validation.Key("body", validation.Required.Error("request body param required"), validation.NilOrNotEmpty.Error("request body param not empty")),
	)); err != nil {
		return errors.New(err.Error()), res
	}

	uri := fmt.Sprintf("%s/%s?head=%s&body=%s", URI, utils.GlobalConfig.Phprm.Token, url.QueryEscape(head), url.QueryEscape(body))
	response, err := resty.New().R().Get(uri)
	if err != nil {
		return errors.New(fmt.Sprintf("请求失败：%s", err.Error())), res
	}

	if code := gjson.Get(response.String(), "code").Int(); code == int64(1) {
		return errors.New(fmt.Sprintf("请求错误：%s", gjson.Get(response.String(), "message").String())), res
	}

	err = json.Unmarshal(response.Body(), &res)
	if err != nil {
		return errors.New("json 格式化数据失败"), res
	}

	return nil, res
}
