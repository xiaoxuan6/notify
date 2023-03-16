package notify_test

import (
    "github.com/stretchr/testify/assert"
    "github.com/xiaoxuan6/notify/v3"
    "github.com/xiaoxuan6/notify/v3/push_plus"
    "github.com/xiaoxuan6/notify/v3/utils"
    "testing"
)

// go test robot_test.go
var config = utils.LoadConfig("./env.yml")

func TestRobotDingTalk(t *testing.T) {
   robot := notify.NewDingTalk(config)
   err := robot.SendText("golang xxxx", []string{}, []string{}, false)
   assert.Nil(t, err)
}

func TestRobotFeishu(t *testing.T) {
    robot := notify.NewFeiShu(config)
    err := robot.SendText("notify xxxx")
    assert.Nil(t, err)
}

// go test -v -run TestRobotPhprm
func TestRobotPhprm(t *testing.T) {
   robot := notify.NewPhprm(config)
   err, r := robot.Send("hello phprm", "这是测试内容")
   assert.Nil(t, err)
   assert.Contains(t, r.Message, "请求成功")
}

func TestRobotPushPlus(t *testing.T) {
   robot := notify.NewPushPlus(config)
   err, res := robot.Send(push_plus.Message{Title: "xxx", Content: "xxxx"})
   assert.Nil(t, err)
   assert.Equal(t, res.Code, 200)
   assert.Contains(t, res.Msg, "请求成功")
}

func TestRobotServer(t *testing.T) {
   robot := notify.NewServer(config)
   _, err := robot.Send("xxx", "xxx")
   assert.Nil(t, err)
}

func TestRobotWechat(t *testing.T) {
   robot := notify.NewWechat(config)
   err := robot.SendText("xxx", []string{}, []string{})
   assert.Nil(t, err)
}

func TestRobotZhiXi(t *testing.T) {
    robot := notify.NewZhiXi(config)
    err, r := robot.Send("ss", "sss")
    assert.Nil(t, err)
    assert.Equal(t, r.Code, 200)
}
