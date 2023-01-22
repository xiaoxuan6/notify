package phprm

import (
	"encoding/json"
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"net/url"
)

type Robot struct {
	Token string `json:"token"`
}

const URI = "https://www.phprm.com/services/push/trigger"

func NewRobot(token string) *Robot {
	return &Robot{
		Token: token,
	}
}

func (r *Robot) Send(head, body string) (err error, res *Response) {

	requestParams := map[string]interface{}{
		"token": r.Token,
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

	var params string
	params = fmt.Sprintf("head=%s&body=%s", url.QueryEscape(head), url.QueryEscape(body))
	uri := fmt.Sprintf("%s/%s?%s", URI, r.Token, params)
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
