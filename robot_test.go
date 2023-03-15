package notify

import (
    "github.com/stretchr/testify/assert"
    "github.com/xiaoxuan6/notify/v3/utils"
    "testing"
)

// go test -v -run TestRobotPhprm
func TestRobotPhprm(t *testing.T) {

    config := utils.LoadConfig("./env.yml")

    robot := NewPhprm(config)
    err, r := robot.Send("hello phprm", "这是测试内容")
    assert.Nil(t, err)
    assert.Contains(t, r.Message, "请求成功")
}

func TestRobotZhiXi(t *testing.T) {
    config := utils.LoadConfig("./env.yml")

    robot := NewZhiXi(config)
    err, r := robot.Send("ss", "sss")
    assert.Nil(t, err)
    assert.Equal(t, r.Code, 200)
}
