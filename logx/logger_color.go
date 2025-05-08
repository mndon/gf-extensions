package logx

import (
	"github.com/fatih/color"
	"github.com/gogf/gf/v2/os/glog"
)

var defaultLevelColor = map[int]int{
	glog.LEVEL_DEBU: glog.COLOR_YELLOW,
	glog.LEVEL_INFO: glog.COLOR_GREEN,
	glog.LEVEL_NOTI: glog.COLOR_CYAN,
	glog.LEVEL_WARN: glog.COLOR_MAGENTA,
	glog.LEVEL_ERRO: glog.COLOR_RED,
	glog.LEVEL_CRIT: glog.COLOR_HI_RED,
	glog.LEVEL_PANI: glog.COLOR_HI_RED,
	glog.LEVEL_FATA: glog.COLOR_HI_RED,
}

// getColoredStr returns a string that is colored by given color.
func getColoredStr(c int, s string) string {
	return color.New(color.Attribute(c)).Sprint(s)
}

func getColorByLevel(level int) int {
	return defaultLevelColor[level]
}
