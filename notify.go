package notify

import (
	"errors"
	"github.com/spf13/viper"
	talk "github.com/xiaoxuan6/ding-talk"
	wechat_talk "github.com/xiaoxuan6/wechat-talk"
	"sync"
)

const (
	DingTalk   = "dinging"
	WechatTalk = "wechat"
)

type fn func() interface{}

var (
	s        sync.Mutex
	adapters = make(map[string]fn)
)

// RegisterAdapter 注册转接器，名称不可重复
func RegisterAdapter(name string, f func() interface{}) {
	s.Lock()
	defer s.Unlock()

	if name == "" {
		panic("RegisterAdapter: adapter must have a name")
	}

	if _, ok := adapters[name]; ok {
		panic("RegisterAdapter: adapter named " + name + " already registered. ")
	}

	adapters[name] = f
}

// DetectAdapter 根据转接器名称获取转接器实例
func DetectAdapter(name string) (fn, error) {
	if adp, ok := adapters[name]; ok {
		return adp, nil
	}

	if len(adapters) == 0 {
		return nil, errors.New("no adapter available")
	}

	if name == "" {
		if len(adapters) == 1 {
			for _, adp := range adapters {
				return adp, nil
			}
		}
		return nil, errors.New("multiple adapters available; must choose one")
	}

	return nil, errors.New("unknown adapter " + name)
}

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("配置文件 config.yml 不存在")
		} else {
			panic("配置文件读取失败")
		}
	}

	RegisterAdapter(DingTalk, func() interface{} {
		robot := talk.NewRobot(viper.GetString("dinging.access_token"))

		if secret := viper.GetString("dinging.secret"); secret != "" {
			robot.SetSecret(secret)
		}

		return robot
	})

	RegisterAdapter(WechatTalk, func() interface{} {
		return wechat_talk.NewRobot(viper.GetString("wechat.key"))
	})
}
