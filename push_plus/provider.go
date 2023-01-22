package push_plus

import "github.com/xiaoxuan6/notify/v2/utils"

func RegisterProvider(config *utils.Config) *Robot {
	return NewRobot(config.PushPlus.Token)
}
