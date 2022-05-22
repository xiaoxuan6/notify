package notify_test

import (
	"github.com/stretchr/testify/assert"
	talk "github.com/xiaoxuan6/ding-talk"
	"github.com/xiaoxuan6/notify"
	wechat_talk "github.com/xiaoxuan6/wechat-talk"
	"testing"
)

func init() {
	notify.Init()
}

func TestDingTalkAdapter(t *testing.T) {
	f, err := notify.DetectAdapter(notify.DingTalk)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	robot := f().(*talk.Robot)
	err = robot.SendText("123", []string{}, []string{}, false)
	assert.Nil(t, err)
}

func TestWechatTalkAdapter(t *testing.T) {
	f, err := notify.DetectAdapter(notify.WechatTalk)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	robot := f().(*wechat_talk.Robot)
	err = robot.SendText("123", []string{}, []string{})
	assert.Nil(t, err)
}
