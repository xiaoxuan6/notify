package zhixi

import (
    "encoding/json"
    "fmt"
    validation "github.com/go-ozzo/ozzo-validation/v4"
    "github.com/go-resty/resty/v2"
    "github.com/pkg/errors"
    "github.com/xiaoxuan6/notify/v3/utils"
)

const URL = "https://xizhi.qqoq.net/%s.send"

type Robot struct {
}

type Response struct {
    Code int    `json:"code"`
    Msg  string `json:"msg"`
}

func (r *Robot) Send(title, content string) (error, Response) {
    var response Response

    requestParams := map[string]interface{}{
        "title":   title,
        "content": content,
    }
    if err := validation.Validate(requestParams, validation.Map(
        validation.Key("title", validation.Required.Error("request title param required"), validation.NilOrNotEmpty.Error("request title param not empty")),
        validation.Key("content", validation.Required.Error("request body content required"), validation.NilOrNotEmpty.Error("request content param not empty")),
    )); err != nil {
        return errors.New(err.Error()), response
    }

    url := fmt.Sprintf(URL, utils.GlobalConfig.ZhiXi.Token)
    res, err := resty.New().R().SetBody(requestParams).Post(url)
    if err != nil {
        return err, response
    }

    if err = json.Unmarshal(res.Body(), &response); err != nil {
        return errors.New(fmt.Sprintf("json 解析错误：%s", err.Error())), response
    }

    return nil, response
}
