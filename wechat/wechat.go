package wechat

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "github.com/go-resty/resty/v2"
    "github.com/xiaoxuan6/notify/v3/utils"
    "io"
    "io/ioutil"
    "mime/multipart"
    "net/http"
    "time"
)

const webhook = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="

const (
    msgTypeText     = "text"
    msgTypeMarkdown = "markdown"
    msgTypeNews     = "news"
    msgTypeImage    = "image"
    msgTypeFile     = "file"
)

type Robot struct {
}

func (r *Robot) SendText(content string, mentionedList, mentionedMobileList []string) error {
    return r.send(&textMessage{
        MsgType: msgTypeText,
        Text: text{
            Content:             content,
            MentionedList:       mentionedList,
            MentionedMobileList: mentionedMobileList,
        },
    })
}

func (r *Robot) SendMarkdown(content string) error {
    return r.send(&markdownMessage{
        MsgType: msgTypeMarkdown,
        Markdown: markdown{
            Content: content,
        },
    })
}

// SendImage 注：图片（base64编码前）最大不能超过2M，支持JPG,PNG格式
func (r *Robot) SendImage(base64, md5 string) error {
    return r.send(&imageMessage{
        MsgType: msgTypeImage,
        Image: image{
            Base64: base64,
            Md5:    md5,
        },
    })
}

func (r *Robot) SendNews(articles []Articles) error {
    return r.send(&newsMessage{
        MsgType: msgTypeNews,
        News: news{
            Articles: articles,
        },
    })
}

func (r *Robot) SendFile(mediaId string) error {
    return r.send(&fileMessage{
        MsgType: msgTypeFile,
        File: file{
            MediaId: mediaId,
        },
    })
}

func (r *Robot) send(msg interface{}) (err error) {

    body, er := json.Marshal(msg)
    if er != nil {
        return errors.New("json 格式化错误")
    }

    uri := fmt.Sprintf("%s%s", webhook, utils.GlobalConfig.Wechat.Key)
    res, err := resty.New().R().
        SetHeader("Content-Type", "application/json;charset=utf-8").
        SetBody(string(body)).
        Post(uri)

    if err != nil {
        return err
    }

    var item = make(map[string]interface{})
    _ = json.Unmarshal(res.Body(), &item)

    if item["errcode"] == float64(0) {
        return nil
    }

    return errors.New(item["errmsg"].(string))
}

//注意client 本身是连接池，不要每次请求时创建client
var (
    HttpClient = &http.Client{
        Timeout: 3 * time.Second,
    }
)

// UploadFile 素材上传得到media_id，该media_id仅三天内有效
// media_id只能是对应上传文件的机器人可以使用
// 要求文件大小在5B~20M之间
func (r *Robot) UploadFile(filename string, file io.Reader) (string, error) {
    url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/upload_media?key=%s&type=file", utils.GlobalConfig.Wechat.Key)

    body := new(bytes.Buffer)
    writer := multipart.NewWriter(body)
    formFile, err := writer.CreateFormFile("media", filename)
    if err != nil {
        return "", err
    }

    _, err = io.Copy(formFile, file)
    if err != nil {
        return "", err
    }

    err = writer.Close()
    if err != nil {
        return "", err
    }

    req, err := http.NewRequest("POST", url, body)
    if err != nil {
        return "", err
    }

    req.Header.Add("Content-Type", writer.FormDataContentType())
    resp, err := HttpClient.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    content, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    data := make(map[string]interface{})
    err = json.Unmarshal(content, &data)
    if err != nil {
        return "", errors.New("json 解析数据失败")
    }

    if data["errcode"] != float64(0) {
        return "", errors.New(data["errmsg"].(string))
    }

    return data["media_id"].(string), nil
}
