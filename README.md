# 消息通知

目前仅支持钉钉机器人和企业微信机器人

# Installation

    go get github.com/xiaoxuan6/notify

# Usage

1、在根目录下创建 `config.yml` 文件，内容：

```yaml
# 钉钉
dinging:
  access_token: "xxx"
  # 设置了 secret 表示对消息内容进行加密
  secret: "xxx"

# 企业微信
wechat:
  key: "xxx"
```

2、初始化配置信息

```bigquery
func init() {
    notify.Init()
}
```

2、使用

```go
f, err := notify.DetectAdapter(name) // name:notify.DingTalk、notify.WechatTalk
if err != nil {
    fmt.Println(err.Error())
}

robot := f().(*talk.Robot) // 这里断言上面适配器取出的是否是对应机器人主体
```