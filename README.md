# 消息通知

目前仅支持钉钉群自定义机器人、企业微信群自定义机器人、飞书群自定义机器人、一封传话、pushplus 和 Server 酱消息通知

# Installation

    go get github.com/xiaoxuan6/notify/v3

# Usage

```php
config := utils.LoadConfig("./env.yml")

robot := notify.NewPhprm(config)
robot.Send("hello phprm", "这是测试内容")
```

`env.yml` 文件在项目根目录，内容如下

```bigquery
dinging:
  access_token:
  secret:
wechat:
  key:
feishu:
  access_token:
  secret:
server:
  webhook:
  secret:
  channel: 9
push_plus:
  token:
phprm:
  token: "xxx"
```