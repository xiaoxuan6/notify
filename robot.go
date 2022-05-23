package notify

import (
	talk "github.com/xiaoxuan6/ding-talk"
	"github.com/xiaoxuan6/notify/server"
	wechat_talk "github.com/xiaoxuan6/wechat-talk"
)

func DingTalkRobot() (robot *talk.Robot, error error) {

	f, err := DetectAdapter(DingTalk)
	if err != nil {
		return nil, err
	}

	robot = f().(*talk.Robot)

	return robot, nil
}

func WechatTalkRobot() (robot *wechat_talk.Robot, error error) {
	f, err := DetectAdapter(WechatTalk)
	if err != nil {
		return nil, err
	}

	robot = f().(*wechat_talk.Robot)

	return robot, nil
}

func ServerRobot() (robot *server.Robot, error error) {
	f, err := DetectAdapter(ServerTalk)

	if err != nil {
		return nil, err
	}

	robot = f().(*server.Robot)

	return robot, nil
}
