package copyx

import (
	"github.com/gogf/gf/v2/util/gconv"
	"reflect"
)

// Copy
// @Description: 复制
// @param fromValue
// @param toValue
// @return T
func Copy[T any](fromValue any, toValue T) T {
	if !reflect.ValueOf(fromValue).IsNil() {
		err := gconv.Scan(fromValue, toValue)
		if err != nil {
			panic(err)
		}
	}
	return toValue
}
