package notify_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/xiaoxuan6/notify/v3"
	"github.com/xiaoxuan6/notify/v3/push_plus"
	"github.com/xiaoxuan6/notify/v3/utils"
	"testing"
)

func TestDingingTalk(t *testing.T) {

	var accessToken = "xxx"
	var secret = "xxx"
	config := `{"dinging":{"access_token":"` + accessToken + `","secret":"` + secret + `"}}`

	con := &utils.Config{}
	_ = json.Unmarshal([]byte(config), con)

	robot := notify.NewNotify(con).DingDing

	err := robot.SendText("test", []string{}, []string{}, false)

	assert.Nil(t, err)
}

func TestFeishuTalk(t *testing.T) {

	var accessToken = "xxx"
	var secret = "xxx"
	config := `{"feishu":{"access_token":"` + accessToken + `","secret":"` + secret + `"}}`

	con := &utils.Config{}
	_ = json.Unmarshal([]byte(config), con)

	robot := notify.NewNotify(con).Feishu

	err := robot.SendText("wer")

	assert.Nil(t, err)
}

func TestServer(t *testing.T) {
	config := `{"server":{"webhook":"https://sctapi.ftqq.com/xxx.send","channel":` + utils.FangtangChannel + `}}`

	con := &utils.Config{}
	_ = json.Unmarshal([]byte(config), con)

	robot := notify.NewNotify(con).Server
	_, err := robot.Send("123", "123")

	assert.Nil(t, err)
}

func TestPushPlus(t *testing.T) {
	var Token = ""
	config := &utils.Config{
		PushPlus: utils.PushPlusConfig{
			Token: Token,
		},
	}

	robot := notify.NewNotify(config).PushPlus
	err, result := robot.Send(push_plus.Message{Title: "test", Content: "test"})
	assert.Nil(t, err)
	assert.Contains(t, result.Msg, "请求成功")
}

func TestPhprm(t *testing.T) {
	var token = "1d8a3e21fac726dbe6da8bc0e463d50fs"
	config := &utils.Config{
		Phprm: utils.PhprmConfig{
			Token: token,
		},
	}
	robot := notify.NewNotify(config).Phprm
	err, response := robot.Send("hello phprm", "这是测试内容")
	assert.Nil(t, err)
	assert.Equal(t, response.Code, 0)
	assert.Contains(t, response.Message, "请求成功")
}
