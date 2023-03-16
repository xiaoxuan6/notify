package push_plus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var robot = newRobot()

func newRobot() *Robot {
    return &Robot{}
}

func TestSendHtml(t *testing.T) {
	message := Message{
		Title:   "test",
		Content: "hello world",
	}

	err, result := robot.Send(message)
	assert.Nil(t, err)
	assert.Contains(t, result.Msg, "请求成功")
}

func TestSendTxt(t *testing.T) {
	message := Message{
		Title:    "test1",
		Content:  `<h1>纯文本内容</h1>`,
		Template: TEMPLATE_TXT,
	}

	err1, result := robot.Send(message)
	assert.Nil(t, err1)
	assert.Contains(t, result.Msg, "请求成功")
}

func TestSendJson(t *testing.T) {
	message := Message{
		Title:    "test1",
		Content:  `{"content", "hello world"}`,
		Template: TEMPLATE_JSON,
	}

	err1, result := robot.Send(message)
	assert.Nil(t, err1)
	assert.Contains(t, result.Msg, "请求成功")
}

func TestSendMarkdown(t *testing.T) {
	message := Message{
		Title:    "test",
		Content:  "# 内容",
		Template: TEMPLATE_MARKDOWN,
	}

	err, result := robot.Send(message)
	assert.Nil(t, err)
	assert.Contains(t, result.Msg, "请求成功")
}

func TestSendWebhook(t *testing.T) {
	message := Message{
		Title:    "test",
		Content:  "# 内容首页 golang",
		Template: TEMPLATE_MARKDOWN,
		Channel:  CHANNEL_WEBHOOK,
		//Webhook:  "dingding", // 钉钉机器人，目前不支持加密模式
		Webhook: "wechat", // 企业微信机器人
	}

	err, result := robot.Send(message)
	assert.Nil(t, err)
	assert.Contains(t, result.Msg, "请求成功")
}
