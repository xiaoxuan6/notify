# 消息通知

目前仅支持钉钉群自定义机器人、企业微信群自定义机器人、飞书群自定义机器人、一封传话、pushplus 和 Server 酱消息通知

# Installation

    go get github.com/xiaoxuan6/notify/v3

# Usage
```go

```

## Example

```php
var token = "1d8a3e21fac726dbe6da8bc0e463d50fs"
config := &utils.Config{
    Phprm: utils.Phprm{
        Token: token,
    },
}
robot := notify.NewPhprm(config)
robot.Send("hello phprm", "这是测试内容")
```
