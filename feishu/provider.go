package feishu

import (
	"github.com/xiaoxuan6/notify/v2/utils"
)

func RegisterProvider(config *utils.Config) *Robot {

	robot := NewRobot(config.Feishu.AccessToken)

	if len(config.Feishu.Secret) > 0 {
		robot.SetSecret(config.Feishu.Secret)
	}

	return robot
}
