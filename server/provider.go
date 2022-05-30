package server

import (
	"github.com/xiaoxuan6/notify/v2/utils"
)

func RegisterProvider(config *utils.Config) *Robot {

	robot := NewRobot(config.Server.Webhook)

	if len(config.Server.Channel) > 0 {
		robot.SetChannel(config.Server.Channel)
	}

	if len(config.Server.Secret) > 0 {
		robot.SetSecret(config.Server.Secret)
	}

	return robot
}
