package logx

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestUse(t *testing.T) {
	// x.Log().Type("type").Info(ctx, "3333")
	// logx.New().Type("type").Info(ctx, "3333")

}

func TestWithCustomFields(t *testing.T) {
	g.Log().SetHandlers(HandlerJson)
	ctx := gctx.New()

	ctx = WithCustomFields(ctx, CustomFields{
		Uid: "uid",
	})

	New(ctx).Type("ACCESS").Info("x0")

	g.Log().Info(ctx, "x1") // 输出无type=ACCESS
	New(ctx).Info("x2")     // 输出无type=ACCESS
}
