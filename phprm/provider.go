package phprm

import "github.com/xiaoxuan6/notify/v3/utils"

func RegisterProvider(config *utils.Config) *Robot {
	return NewRobot(config.Phprm.Token)
}
