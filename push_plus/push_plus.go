package push_plus

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

const URI = "https://www.pushplus.plus/send"

type Root struct {
	token string
}

func NewRoot(toke string) *Root {
	return &Root{
		token: toke,
	}
}

func (r *Root) Send(message Message) (error error, response *Response) {
	res := &Response{}
	message.Token = r.token

	if len(r.token) == 0 {
		return errors.New("invalid token"), res
	}

	result, err := resty.New().R().SetBody(message).SetHeader("Content-Type", "application/json").Post(URI)
	fmt.Println(result)
	if err != nil {
		return errors.New("json 格式化数据失败"), res
	}
	json.Unmarshal(result.Body(), res)
	if res.Code == 200 {
		return nil, res
	}

	return errors.New(fmt.Sprintf("%s", res.Msg)), res
}
