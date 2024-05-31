package base

import (
	"context"
	"math/rand"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	ctx := context.Background() // 链路起点或者调用的起点
	ctx = context.WithValue(ctx, "key", "value")
	val, ok := ctx.Value("key").(string)
	if !ok {
		t.Log("类型不对")
		return
	}
	t.Log(val, ctx.Done(), ctx.Err())
	t.Log(ctx.Deadline())

	ctx = context.TODO() // 不确定 context 的作用
	ctx = context.WithValue(ctx, "key", "value")
	val, ok = ctx.Value("key").(string)
	if !ok {
		t.Log("类型不对")
		return
	}
	t.Log(<-ctx.Done())
	t.Log(ctx.Err())
	t.Log(ctx.Deadline())
}

func TestWithCancel(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		time.Sleep(time.Second)
		cancel()
	}()

	t.Log(<-ctx.Done())
	t.Log(ctx.Err())
	t.Log(ctx.Deadline())
}

func TestWithDeadline(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
	defer cancel()

	deadline, ok := ctx.Deadline()
	t.Log(deadline, ok)

	go func() {
		if rand.Intn(7) > 3 {
			cancel()
		}
	}()

	t.Log(<-ctx.Done())
	t.Log(ctx.Err())
	t.Log(ctx.Deadline())

}

func TestWithTimeOut(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	deadline, ok := ctx.Deadline()
	t.Log(deadline, ok)

	go func() {
		if rand.Intn(7) > 3 {
			cancel()
		}
	}()

	t.Log(<-ctx.Done())
	t.Log(ctx.Err())
	t.Log(ctx.Deadline())
}

func TestWithErr(t *testing.T) {
	t.Log(context.Canceled)
	t.Log(context.DeadlineExceeded)
}

func TestParentSameKey(t *testing.T) {
	ctx := context.Background()
	parent := context.WithValue(ctx, "my-key", "parent value")
	pVal, ok := parent.Value("my-key").(string)
	t.Log(pVal, ok)

	child := context.WithValue(parent, "my-key", "child value")
	cVal, ok := child.Value("my-key").(string)
	t.Log(cVal, ok)

}

func TestParentTimeoutVal(t *testing.T) {
	ctx := context.Background()
	parent := context.WithValue(ctx, "my-key", "parent value")
	child2, cancel := context.WithTimeout(parent, time.Second)
	defer cancel()
	t.Log(child2.Value("my-key"))
}

func TestParentNotGetChildrenVal(t *testing.T) {
	ctx := context.Background()
	parent := context.WithValue(ctx, "my-key", "parent value")
	child := context.WithValue(parent, "new-key", "child value")
	t.Log(parent.Value("new-key"))
	t.Log(child.Value("new-key"))
}

func TestParentGetChildrenVal(t *testing.T) {
	ctx := context.Background()
	parent := context.WithValue(ctx, "map", map[string]string{})
	child, cancel := context.WithTimeout(parent, time.Second)
	defer cancel()

	cm := child.Value("map").(map[string]string)
	cm["key1"] = "value1"

	pm := parent.Value("map").(map[string]string)
	t.Log(pm["key1"])
}

func TestChildNotSetParentTimeout(t *testing.T) {
	ctx := context.Background()
	parent, pCancel := context.WithTimeout(ctx, time.Second)
	defer pCancel()

	child, cCancel := context.WithTimeout(parent, time.Second*3)
	defer cCancel()

	go func() {
		t.Log(<-child.Done())
		t.Log(child.Err())
		t.Log(child.Deadline())
		t.Log("收到了结束信号")
	}()

	time.Sleep(time.Second * 2)
}

func TestChildSetParentTimeout(t *testing.T) {
	ctx := context.Background()
	parent, pCancel := context.WithTimeout(ctx, time.Second*3)
	defer pCancel()
	child, cCancel := context.WithTimeout(parent, time.Second)
	defer cCancel()

	go func() {
		<-child.Done()
		t.Log("收到了结束信号")
	}()
	time.Sleep(time.Second * 2)
}

func TestTimeoutBiz(t *testing.T) {
	n := rand.Intn(5)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(n))
	defer cancel()
	biz := func() {
		time.Sleep(time.Second * 2)
	}

	ch := make(chan struct{})
	go func() {
		biz()
		ch <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		t.Log("超时了")
	case <-ch:
		t.Log("业务正常结束")
	}
}
