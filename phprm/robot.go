package phprm

import (
	"encoding/json"
	"errors"
	"fmt"
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
	if len(r.Token) == 0 {
		return errors.New("invalid token"), res
	}

	var params string
	if len(head) > 0 {
		head = url.QueryEscape(head)
		params = fmt.Sprintf("head=%s", head)
	} else {
		return errors.New("request head param not empty"), res
	}

	if len(body) > 0 {
		body = url.QueryEscape(body)
		params = fmt.Sprintf("%s&body=%s", params, body)
	} else {
		return errors.New("request body param not empty"), res
	}

	uri := fmt.Sprintf("%s/%s?%s", URI, r.Token, params)
	response, err := resty.New().R().Get(uri)
	if err != nil {
		return errors.New(fmt.Sprintf("请求失败：%s", err.Error())), res
	}

	if code := gjson.Get(response.String(), "Code").Int(); code != 0 {
		return errors.New(fmt.Sprintf("请求错误：%s", gjson.Get(response.String(), "message").String())), res
	}

	err = json.Unmarshal(response.Body(), &res)
	if err != nil {
		return errors.New("json 格式化数据失败"), res
	}

	fmt.Println(res)

	return nil, res
}
