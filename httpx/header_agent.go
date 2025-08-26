package httpx

import (
	"context"
	"regexp"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

const ctxKeyForAgent = "httpxAgent"

type Agent struct {
	agent string

	channel    string // 渠道
	platform   string // 平台
	appVersion string // app版本
	brand      string // 设备厂商
	model      string // 设备型号
	passport   string // 设备id
}

func NewAgent(agent string) *Agent {
	return &Agent{agent: agent}
}

func AgentFromCtx(ctx context.Context) (agent *Agent) {
	v := ctx.Value(ctxKeyForAgent)
	if v != nil {
		return v.(*Agent)
	}

	agent = &Agent{}
	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		return
	}
	if r.Header.Get(HeaderXA) != "" {
		agent.agent = r.Header.Get(HeaderXA)
	} else {
		agent.agent = r.Header.Get(HeaderUA)
	}
	return agent
}

func AgentStrFromHeader(ctx context.Context) (agent string) {
	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		return
	}
	if r.Header.Get(HeaderXA) != "" {
		return r.Header.Get(HeaderXA)
	}
	return r.Header.Get(HeaderUA)
}

func WithAgent(ctx context.Context, agent *Agent) context.Context {
	return context.WithValue(ctx, ctxKeyForAgent, agent)
}

func (a *Agent) GetAgent() string {
	return a.agent
}

func (a *Agent) SetChannel(v string) {
	a.channel = strings.ToUpper(v)
}

func (a *Agent) GetChannel() string {
	a.channel = a.GetValueFormUA(`Channel\((.*?)\)`)
	return a.channel
}

func (a *Agent) SetPlatform(v string) {
	a.platform = strings.ToUpper(v)
}

func (a *Agent) GetPlatform() string {
	if a.platform == "" {
		a.platform = a.GetValueFormUA(`Platfo.?m\((.*?)\)`)
	}
	return a.platform
}

func (a *Agent) SetAppVersion(v string) {
	a.appVersion = strings.ToUpper(v)
}

func (a *Agent) GetAppVersion() string {
	if a.appVersion == "" {
		a.appVersion = a.GetValueFormUA(`App Version\((.*?)\)`)
	}
	return a.appVersion
}

func (a *Agent) SetBrand(v string) {
	a.brand = strings.ToUpper(v)
}

func (a *Agent) GetBrand() string {
	if a.brand == "" {
		a.brand = a.GetValueFormUA(`Brand\((.*?)\)`)
	}
	return a.brand
}

func (a *Agent) SetModel(v string) {
	a.model = strings.ToUpper(v)
}

func (a *Agent) GetModel() string {
	if a.model == "" {
		a.model = a.GetValueFormUA(`Model\((.*?)\)`)
	}
	return a.model
}

func (a *Agent) SetPassport(v string) {
	a.passport = strings.ToUpper(v)
}

func (a *Agent) GetPassport() string {
	if a.passport == "" {
		a.passport = a.GetValueFormUA(`Passport\((.*?)\)`)
	}
	return a.passport
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
