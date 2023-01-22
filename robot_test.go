package notify

import (
	"github.com/stretchr/testify/assert"
	"github.com/xiaoxuan6/notify/v2/utils"
	"testing"
)

func TestRobotPhprm(t *testing.T) {
	var token = "1d8a3e21fac726dbe6da8bc0e463d50fs"
	config := &utils.Config{
		Phprm: utils.Phprm{
			Token: token,
		},
	}
	root := NewPhprm(config)
	err, r := root.Send("hello phprm", "这是测试内容")
	assert.Nil(t, err, err.Error())
	assert.Contains(t, r.Message, "请求成功")
}
