package http

import (
	"context"
	"github.com/mndon/gf-extensions/config"
)

const (
	zh = "zh-CN"
	en = "en"
)

var local string

func Init() {
	local = config.GetWithCmdFromCfgWithPanic(context.TODO(), "local", en).String()
}

type I18nMsg struct {
	Zh string
	En string
}

func (i *I18nMsg) GetMsg() string {
	if local == "" {
		Init()
	}

	if local == zh {
		return i.Zh
	} else {
		return i.En
	}
}
