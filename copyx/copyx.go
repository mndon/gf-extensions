package copyx

import (
	"github.com/gogf/gf/v2/util/gconv"
)

// Copy *struct => *struct or *struct => **struct
func Copy[T any](fromValue interface{}, toValue T) T {
	err := gconv.Scan(fromValue, toValue)
	if err != nil {
		panic(err)
	}
	return toValue
}
