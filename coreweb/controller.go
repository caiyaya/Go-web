package main

import (
	"context"
	"coreweb/framework"
	"fmt"
	"log"
	"time"
)

func FooControllerHandler(c *framework.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Second))
	defer cancel()

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// do real action
		time.Sleep(10 * time.Second)
		c.Json(200, "ok")
		finish <- struct{}{}
	}()
	// 加锁目的是保证 在 调用 c.Json() 过程中不会对responseWriter造成并发写入的乱序（主要和业务代码的冲突，超时了，但是业务代码的协程依然在进行）|| 重写 增加tag白哦职位
	select {
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.Json(500, "panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.Json(500, "time out")
		c.SetHasTimeout()
	}
	return nil
}
