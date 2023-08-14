package copyx

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/jinzhu/copier"
	"reflect"
)

type Option struct {
	IgnoreEmpty bool
	DeepCopy    bool
}

func createNewVariable(ptr interface{}) error {
	ptrType := reflect.TypeOf(ptr)
	ptrValue := reflect.ValueOf(ptr)
	// 空值判断
	if ptrType == nil {
		return gerror.NewCode(gcode.CodeInvalidParameter, "parameter pointer should not be nil")
	}
	// **struct处理
	if ptrValue.Kind() == reflect.Ptr && ptrValue.Elem().Kind() == reflect.Ptr {
		if !ptrValue.Elem().Elem().IsValid() {
			newVar := reflect.New(ptrType.Elem().Elem()).Elem()
			ptrValue.Elem().Set(newVar.Addr())
		}

	}
	return nil
}

// Copy *struct => *struct or *struct => **struct
func Copy(fromValue interface{}, toValue interface{}) (err error) {
	err = createNewVariable(toValue)
	if err != nil {
		return err
	}
	return CopyWithOption(fromValue, toValue, Option{
		IgnoreEmpty: true,
		DeepCopy:    false,
	})
}

// CopyWithOption copy with option
func CopyWithOption(fromValue interface{}, toValue interface{}, opt Option) (err error) {
	return copier.CopyWithOption(toValue, fromValue, copier.Option{
		IgnoreEmpty: opt.IgnoreEmpty,
		DeepCopy:    opt.DeepCopy,
	})
}
