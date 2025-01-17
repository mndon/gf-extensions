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

	g.Log().SetHandlers(HandlerLocal)
	//g.Log().SetHandlers(HandlerJson)z
	err := gerror.New("fdsafsdf")
	New().Type("111").Errorf(ctx, "%+v", err)
	New().Type("111").Infof(ctx, "%+v", err)
	New().Type("111").Infof(ctx, "ccccc")
}
