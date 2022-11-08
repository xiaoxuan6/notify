# 消息通知

目前仅支持钉钉群自定义机器人、企业微信群自定义机器人、飞书群自定义机器人、pushplus 和 Server 酱消息通知

# Installation

    go get github.com/xiaoxuan6/notify/v2

# Usage
```go
config := `{"server":{"webhook":"https://sctapi.ftqq.com/xxx.send","channel":` + utils.FangtangChannel + `}}`

con := &utils.Config{}
_ = json.Unmarshal([]byte(config), con)

robot := notify.NewNotify(con).Server
_, err := robot.Send("123", "123")
```

OR
```go
var Token = ""
config := &utils.Config{
    PushPlus: utils.PushPlusConfig{
        Token: Token,
    },
}

root := notify.NewNotify(config).PushPlus
err, result := root.Send(push_plus.Message{Title: "test", Content: "test"})
```
