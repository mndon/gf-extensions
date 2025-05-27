package delayqueuex

import (
	"github.com/gogf/gf/v2/os/gctx"

	"context"
	"fmt"
	"testing"
	"time"
)

func Handler1(ctx context.Context, r *Message) error {
	fmt.Println("Handler1 called:", r)
	return nil
}

func TestServer_Run(t *testing.T) {
	Register("topic1", Handler1)
	Run()

	id, err := AddTask(gctx.New(), "topic1", "{\"a\": 1}", 1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(id)

	time.Sleep(100 * time.Second)

}
