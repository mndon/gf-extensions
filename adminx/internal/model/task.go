/*
* @desc:定时任务
* @company:
* @Author:
* @Date:   2023/1/13 17:47
 */

package model

import "context"

type TimeTask struct {
	FuncName string
	Param    []string
	Run      func(ctx context.Context)
}
