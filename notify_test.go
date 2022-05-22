package notify_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/xiaoxuan6/notify"
	"github.com/xiaoxuan6/notify/server"
	"testing"
)

func init() {
	notify.Init()
}

func TestDingTalkAdapter(t *testing.T) {
	robot, err := notify.DingTalkRobot()
	assert.Nil(t, err)

	err = robot.SendText("123", []string{}, []string{}, false)
	assert.Nil(t, err)
}

func TestWechatTalkAdapter(t *testing.T) {
	robot, err := notify.WechatTalkRobot()
	assert.Nil(t, err)

	err = robot.SendText("123", []string{}, []string{})
	assert.Nil(t, err)
}

func TestServerRobotAdapter(t *testing.T) {
	robot, err := notify.ServerRobot()
	assert.Nil(t, err)

	resp, err := robot.Send("123", "32423")
	assert.Nil(t, err)

	assert.Equal(t, float64(0), resp["errno"])
}

func TestServerRobotWithChannel1Adapter(t *testing.T) {
	_ = notify.DeleteAdapter(notify.ServerTalk)

	notify.RegisterAdapter(notify.ServerTalk, func() interface{} {
		return server.NewRobot("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxx").SetChannel("1")
	})

	robot, _ := notify.ServerRobot()
	_, err := robot.Send("123", "32423")
	assert.Nil(t, err)
}

func TestServerRobotWithChannel2Adapter(t *testing.T) {
	_ = notify.DeleteAdapter(notify.ServerTalk)

	notify.RegisterAdapter(notify.ServerTalk, func() interface{} {
		webhook := "https://oapi.dingtalk.com/robot/send?access_token=xxx&server=xxx"
		return server.NewRobot(webhook).SetChannel("2")
	})

	robot, _ := notify.ServerRobot()
	_, err := robot.Send("123", "32423")
	assert.Nil(t, err)
}
