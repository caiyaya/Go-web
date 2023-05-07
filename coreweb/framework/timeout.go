package framework

import (
	"context"
	"fmt"
	"log"
	"time"
)

func TimeoutHandler(fun ControllerHandler, d time.Duration) ControllerHandler {
	// 使用函数回调
	return func(c *Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		// 执行业务逻辑前 预操作 初始化 超时 context
		durationCtx, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()

		c.request.WithContext(durationCtx)
		// 执行业务操作
		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			// 执行具体的业务逻辑
			fun(c)
			finish <- struct{}{}
		}()

		// 执行业务逻辑后操作
		select {
		case p := <-panicChan:
			log.Println(p)
			c.responseWriter.WriteHeader(500)
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			c.SetHasTimeout()
			c.responseWriter.Write([]byte("time out"))
		}
		return nil
	}
}
