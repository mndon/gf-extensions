package utils

import "time"

// ConvTZLocalTime 将标准时间转化成对应时区的时间
func ConvTZLocalTime(t time.Time, zone time.Duration) time.Time {
	_, offset := t.Zone()
	return t.Add((zone - time.Duration(offset)) * time.Second)
}
