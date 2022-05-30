package notify_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/xiaoxuan6/notify/v2"
	"github.com/xiaoxuan6/notify/v2/utils"
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
