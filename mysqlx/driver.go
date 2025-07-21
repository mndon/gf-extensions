package mysqlx

import (
	"context"
	"fmt"
	"github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mndon/gf-extensions/logx"
)

// 自定义mysql驱动，增强慢查询日志记录，参考：https://goframe.org/docs/core/gdb-interface-callback

type Driver struct {
	*mysql.Driver
}

const (
	transactionIdForLoggerCtx = "TransactionId" //
	// customDriverName is my driver name, which is used for registering.
	customDriverName = "mysql"
)

func init() {
	if err := gdb.Register(customDriverName, &Driver{}); err != nil {
		panic(err)
	}
}

func (d *Driver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error) {
	return &Driver{
		&mysql.Driver{
			Core: core,
		},
	}, nil
}

func (d *Driver) DoCommit(ctx context.Context, in gdb.DoCommitInput) (out gdb.DoCommitOutput, err error) {
	tsMilliStart := gtime.TimestampMilli()
	out, err = d.Core.DoCommit(ctx, in)
	if err != nil {
		return out, err
	}
	tsMilliFinished := gtime.TimestampMilli()
	tsMilliCost := tsMilliFinished - tsMilliStart
	// 获取慢查阈值
	slowExecuteDuration := g.Cfg().MustGet(ctx, "database.logger.slowExecuteDuration", 1000).Int64()
	// 慢查告警
	if tsMilliCost >= slowExecuteDuration {
		var transactionIdStr string
		if in.IsTransaction {
			if v := ctx.Value(transactionIdForLoggerCtx); v != nil {
				transactionIdStr = fmt.Sprintf(`[txid:%d] `, v.(uint64))
			}
		}
		var rowsAffected int64
		if out.Result != nil {
			rowsAffected, _ = out.Result.RowsAffected()
		}
		s := fmt.Sprintf(
			"[%3d ms] [%s] [%s] [rows:%-3d] %s%s",
			tsMilliFinished-tsMilliStart, d.GetGroup(), d.GetSchema(), rowsAffected, transactionIdStr, gdb.FormatSqlWithArgs(in.Sql, in.Args),
		)
		logx.New(ctx).Type("mysql_slow_execute").Warning(s)
		// todo 发送告警
	}
	return
}
