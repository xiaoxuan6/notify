package notify

import (
    "github.com/stretchr/testify/assert"
    "github.com/xiaoxuan6/notify/v3/push_plus"
    "github.com/xiaoxuan6/notify/v3/utils"
    "testing"
)

var config = utils.LoadConfig("./env.yml")

//func TestRobotFeishu(t *testing.T) {
//    robot := NewFeiShu(config)
//    err := robot.SendText("xxxx")
//    assert.Nil(t, err)
//}

// go test -v -run TestRobotPhprm
func TestRobotPhprm(t *testing.T) {
    robot := NewPhprm(config)
    err, r := robot.Send("hello phprm", "这是测试内容")
    assert.Nil(t, err)
    assert.Contains(t, r.Message, "请求成功")
}

func TestRobotPushPlus(t *testing.T) {
    robot := NewPushPlus(config)
    err, res := robot.Send(push_plus.Message{Title: "xxx", Content: "xxxx"})
    assert.Nil(t, err)
    assert.Equal(t, res.Code, 200)
    assert.Contains(t, res.Msg, "请求成功")
}

func TestRobotWechat(t *testing.T) {
    robot := NewWechat(config)
    err := robot.SendText("xxx", []string{}, []string{})
    assert.Nil(t, err)
}

func TestRobotZhiXi(t *testing.T) {
    robot := NewZhiXi(config)
    err, r := robot.Send("ss", "sss")
    assert.Nil(t, err)
    assert.Equal(t, r.Code, 200)
}
