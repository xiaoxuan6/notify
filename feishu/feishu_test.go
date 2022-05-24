package feishu

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var robot = NewRobot("xxx")

func TestSendTextAndHash26(t *testing.T) {
	err := robot.SetSecret("xxx").SendText("test")
	assert.Nil(t, err)

	// @所有人
	err = robot.SendText("<at user_id=\"all\">所有人</at> test")
	assert.Nil(t, err)
}

func TestSendText(t *testing.T) {
	err := robot.SendText("test")
	assert.Nil(t, err)

	// @所有人
	err = robot.SendText("<at user_id=\"all\">所有人</at> test")
	assert.Nil(t, err)
}

func TestSendPost(t *testing.T) {
	// text
	content := make([][]ContentParams, 0)
	contentSub := make([]ContentParams, 0)
	contentSub = append(contentSub, ContentParams{
		Tag:  "text",
		Text: "项目更新",
	})
	content = append(content, contentSub)

	err := robot.SendPost("test post", content)
	assert.Nil(t, err)

	// href
	content1 := make([][]ContentParams, 0)
	contentSub1 := make([]ContentParams, 0)
	contentSub1 = append(contentSub1, ContentParams{
		Tag:  "a",
		Text: "请查看",
		Href: "http://www.example.com/",
	})
	content1 = append(content1, contentSub1)

	err = robot.SendPost("test post", content1)
	assert.Nil(t, err)

	// at
	content2 := make([][]ContentParams, 0)
	contentSub2 := make([]ContentParams, 0)
	contentSub2 = append(contentSub2, ContentParams{
		Tag:    "at",
		UserId: "all", // 支持所有人和某个人的 open_id
	})
	content2 = append(content2, contentSub2)

	err = robot.SendPost("test post", content2)
	assert.Nil(t, err)

	// all
	content3 := make([][]ContentParams, 0)
	contentSub3 := make([]ContentParams, 0)
	contentSub3 = append(contentSub3, ContentParams{
		Tag:  "text",
		Text: "项目更新",
	}, ContentParams{
		Tag:  "a",
		Text: "请查看",
		Href: "http://www.example.com/",
	}, ContentParams{
		Tag:    "at",
		UserId: "all",
	})
	content3 = append(content3, contentSub3)

	err = robot.SendPost("test post", content3)
	assert.Nil(t, err)
}

func TestSendShareChat(t *testing.T) {
	err := robot.SendShareChat("oc_f5b1a7eb27ae2c7b6adc2a74faf339ff")
	if err != nil {
		assert.Error(t, err, err.Error())
	}

	assert.Nil(t, err)
}

func TestSendImage(t *testing.T) {
	err := robot.SendImage("img_ecffc3b9-8f14-400f-a014-05eca1a4310g")
	if err != nil {
		assert.Error(t, err, err.Error())
	}

	assert.Nil(t, err)
}
