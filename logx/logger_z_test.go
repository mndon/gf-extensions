package logx

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestUse(t *testing.T) {
	// x.Log().Type("type").Info(ctx, "3333")
	// logx.New().Type("type").Info(ctx, "3333")

}

func TestName(t *testing.T) {
	ctx := gctx.New()

	g.Log().SetHandlers(HandlerPlain)
	//g.Log().SetHandlers(HandlerJson)z
	ctx = WithCustomFields(ctx, CustomFields{Uid: "10086"})
	err := gerror.New("fdsafsdf")
	New(ctx).Type("typ1").Errorf("%+v", err)
	New(ctx).Type("typ2").Infof("%+v", err)
	New(ctx).Type("typ3").Infof("ccccc")
}
