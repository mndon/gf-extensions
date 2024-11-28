package logx

import (
	"context"
	"testing"
)

func TestRotate(t *testing.T) {

	r := NewRotate("./log/", 2, "24h")
	r.RotateChecksTimely(context.TODO())
}
