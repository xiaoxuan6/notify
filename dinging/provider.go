package dinging

import (
	dinging_talk "github.com/xiaoxuan6/ding-talk"
	"github.com/xiaoxuan6/notify/v2/utils"
)

func RegisterProvider(config *utils.Config) *dinging_talk.Robot {

	robot := dinging_talk.NewRobot(config.Dinging.AccessToken)

	if len(config.Dinging.Secret) > 0 {
		robot.SetSecret(config.Dinging.Secret)
	}

	return robot
}
