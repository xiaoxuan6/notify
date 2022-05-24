# 消息通知

目前仅支持钉钉群自定义机器人、企业微信群自定义机器人、飞书群自定义机器人和 Server 酱消息通知

# Installation

    go get github.com/xiaoxuan6/notify

# Usage

1、将 `config.yml` 复制到项目根目录下，并修改对应的配置信息

2、初始化配置信息

```bigquery
func init() {
    notify.Init()
}
```

2、使用

```go
robot, err := notify.WechatTalkRobot()
if err != nil {
fmt.Println(err.Error())
}
```
