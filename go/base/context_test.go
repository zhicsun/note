package base

import (
	"context"
	"math/rand"
	"testing"
	"time"
)

func TestBackgroundTodoWithVal(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "value")
	val, ok := ctx.Value("key").(string)
	t.Log(val, ok) // value true

	ctx = context.TODO() // 不确定 context 的作用
	ctx = context.WithValue(ctx, "key", "value")
	val, ok = ctx.Value("key").(string)
	t.Log(val, ok) // value true
}

func TestWithCancel(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		time.Sleep(time.Second)
		cancel()
	}()

	t.Log(<-ctx.Done())   // {}
	t.Log(ctx.Err())      // context canceled
	t.Log(ctx.Deadline()) //  0001-01-01 00:00:00 +0000 UTC false
}

func TestWithDeadline(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
	defer cancel()

	deadline, ok := ctx.Deadline()
	t.Log(deadline, ok) // 2024-06-19 22:48:59.808979 +0800 CST m=+3.000902167 true
	t.Log(ctx.Err())    // nil

	go func() {
		if rand.Intn(7) > 3 {
			cancel()
		}
	}()

	t.Log(<-ctx.Done())   // {}
	t.Log(ctx.Err())      // context deadline exceeded
	t.Log(ctx.Deadline()) // 2024-06-19 22:48:59.808979 +0800 CST m=+3.000902167 true
}

func TestWithTimeOut(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	deadline, ok := ctx.Deadline()
	t.Log(deadline, ok) // 2024-06-19 22:48:59.808979 +0800 CST m=+3.000902167 true
	t.Log(ctx.Err())    // nil

	go func() {
		if rand.Intn(7) > 3 {
			cancel()
		}
	}()

	t.Log(<-ctx.Done())   // {}
	t.Log(ctx.Err())      // context deadline exceeded
	t.Log(ctx.Deadline()) // 2024-06-19 22:48:59.808979 +0800 CST m=+3.000902167 true
}

func TestParentSameKey(t *testing.T) {
	ctx := context.Background()
	parent := context.WithValue(ctx, "key", "parent")

	pVal, ok := parent.Value("key").(string)
	t.Log(pVal, ok) // parent true

	child := context.WithValue(parent, "key", "child")
	cVal, ok := child.Value("key").(string)
	t.Log(cVal, ok) // child true
}

func TestChildTimeoutVal(t *testing.T) {
	ctx := context.Background()
	parent := context.WithValue(ctx, "key", "parent")

	child, cancel := context.WithTimeout(parent, time.Second)
	defer cancel()

	t.Log(<-child.Done())   // {}
	t.Log(child.Err())      // context deadline exceeded
	t.Log(child.Deadline()) // 2024-06-19 22:48:59.808979 +0800 CST m=+3.000902167 true

	val, ok := child.Value("key").(string)
	t.Log(val, ok) // parent true
}

func TestParentNotGetChildrenVal(t *testing.T) {
	ctx := context.Background()
	parent := context.WithValue(ctx, "key", "parent value")
	child := context.WithValue(parent, "cKey", "child value")

	pVal, ok := parent.Value("key").(string)
	t.Log(pVal, ok) // parent value true

	cVal, ok := child.Value("cKey").(string)
	t.Log(cVal, ok) // child value true

	pVal, ok = parent.Value("cKey").(string)
	t.Log(pVal, ok) // "" false
}

func TestParentGetChildrenVal(t *testing.T) {
	ctx := context.Background()
	parent := context.WithValue(ctx, "map", map[string]string{})
	child, cancel := context.WithTimeout(parent, time.Second)
	defer cancel()

	t.Log(<-child.Done())   // {}
	t.Log(child.Err())      // context deadline exceeded
	t.Log(child.Deadline()) // 2024-06-19 22:48:59.808979 +0800 CST m=+3.000902167 true

	cMap, ok := child.Value("map").(map[string]string)
	t.Log(cMap, ok) // map[] true
	cMap["key"] = "value"

	pMap, ok := parent.Value("map").(map[string]string)
	t.Log(pMap, ok) // map[key:value] true
}

func TestChildNotSetParentTimeout(t *testing.T) {
	ctx := context.Background()
	parent, pCancel := context.WithTimeout(ctx, time.Second)
	defer pCancel()

	child, cCancel := context.WithTimeout(parent, time.Second*3)
	defer cCancel()

	go func() {
		t.Log(<-child.Done())   // {}
		t.Log(child.Err())      // context deadline exceeded
		t.Log(child.Deadline()) // 2024-06-19 22:48:59.808979 +0800 CST m=+3.000902167 true
		t.Log("收到了结束信号")
	}()

	time.Sleep(time.Second * 2) // 父取消子也会取消
}

func TestChildSetParentTimeout(t *testing.T) {
	ctx := context.Background()
	parent, pCancel := context.WithTimeout(ctx, time.Second*3)
	defer pCancel()
	child, cCancel := context.WithTimeout(parent, time.Second)
	defer cCancel()

	go func() {
		t.Log(<-parent.Done())
		t.Log(<-child.Done())
		t.Log("收到了结束信号")
	}()

	time.Sleep(time.Second * 2) // 子取消父不会取消
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
