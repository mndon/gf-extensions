package logx

import (
	"context"
	"github.com/gofrs/flock"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtimer"
	"path/filepath"
	"time"
)

// Rotate
// @Description: 日志滚动特性
type Rotate struct {
	path       string
	countLimit int
	interval   time.Duration
}

func NewRotate(path string, countLimit int, intervalStr string) *Rotate {
	interval, err := time.ParseDuration(intervalStr)
	if err != nil {
		g.Log().Errorf(context.TODO(), " parsing duration err: %+v", err)
		interval = time.Hour
	}

	return &Rotate{
		path:       path,
		countLimit: countLimit,
		interval:   interval,
	}
}

// RotateChecksTimely
// @Description: 执行日志滚动检测任务，并携带文件锁
// @receiver r
// @param ctx
func (r *Rotate) RotateChecksTimely(ctx context.Context) {

	defer gtimer.AddOnce(ctx, r.interval, r.RotateChecksTimely)

	fileLock := flock.New(filepath.Join(r.path, "go-lock.lock"))
	locked, err := fileLock.TryLock()
	if err != nil {
		g.Log().Warningf(ctx, "log rotate get file lock err, %+v", err)
	}
	if !locked {
		g.Log().Warningf(ctx, "log rotate get file lock fail, already lock")
	}

	if locked {
		defer fileLock.Unlock()
		r.rotateChecks(ctx)
	}
}

// rotateChecks
// @Description: 日志文件数量检测
// @receiver r
// @param ctx
func (r *Rotate) rotateChecks(ctx context.Context) {
	// 搜索日志文件
	files, err := gfile.ScanDirFile(r.path, "*.log", true)
	if err != nil {
		g.Log().Errorf(ctx, `%+v`, err)
	}
	g.Log().Infof(ctx, "logging rotation start checks: %+v", files)

	// 清除超过限制的日志文件
	backupFiles := garray.NewSortedArray(func(a, b interface{}) int {
		// Sorted by rotated/backup file mtime.
		// The older rotated/backup file is put in the head of array.
		var (
			file1  = a.(string)
			file2  = b.(string)
			result = gfile.MTimestampMilli(file1) - gfile.MTimestampMilli(file2)
		)
		if result <= 0 {
			return -1
		}
		return 1
	})

	for _, file := range files {
		backupFiles.Add(file)
	}

	g.Log().Infof(ctx, `calculated backup files array: %+v`, backupFiles)

	diff := backupFiles.Len() - r.countLimit
	for i := 0; i < diff; i++ {
		path, _ := backupFiles.PopLeft()
		g.Log().Infof(ctx, `remove exceeded backup limit file: %s`, path)
		if err := gfile.Remove(path.(string)); err != nil {
			g.Log().Errorf(ctx, `%+v`, err)
		}
	}
}
