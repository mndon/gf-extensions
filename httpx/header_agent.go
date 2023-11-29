package httpx

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"regexp"
	"strings"
)

type Agent struct {
	agent string
}

func GetAgent(ctx context.Context) (agent Agent) {
	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		return
	}
	if r.Header.Get(HeaderXA) != "" {
		agent.agent = r.Header.Get(HeaderXA)
	} else {
		agent.agent = r.Header.Get(HeaderUA)
	}
	return
}

func (a *Agent) GetAgent() string {
	return a.agent
}

func (a *Agent) GetChannel() string {
	return a.GetValueFormUA(`Channel\((.*?)\)`)
}

func (a *Agent) GetPlatform() string {
	// 兼容早期字段单词错误Platfom
	return a.GetValueFormUA(`Platfo.?m\((.*?)\)`)
}

func (a *Agent) GetAppVersion() string {
	return a.GetValueFormUA(`App Version\((.*?)\)`)
}

func (a *Agent) GetBrand() string {
	return a.GetValueFormUA(`Brand\((.*?)\)`)
}

func (a *Agent) GetModel() string {
	return a.GetValueFormUA(`Model\((.*?)\)`)
}

func (a *Agent) GetPassport() string {
	return a.GetValueFormUA(`Passport\((.*?)\)`)
}

func (a *Agent) GetValueFormUA(matchKey string) string {
	if a.agent == "" {
		return ""
	}

	brandReg := regexp.MustCompile(matchKey)
	result := brandReg.FindStringSubmatch(a.agent)
	if result != nil && len(result) == 2 {
		return strings.ToUpper(result[1])
	}
	return ""
}
