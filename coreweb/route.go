package main

import (
	"coreweb/framework"
	"time"
)

//func registerRouter(core *framework.Core) {
//	// 设置控制器
//	core.Get("foo", FooControllerHandler)
//}

// 注册路由
func registerRouter(core *framework.Core) {
	// http方法 + 静态路由匹配
	//core.Get("user/login", UserLoginController)
	// 在核心业务逻辑 UserLoginController 之外，封装一层 TimeoutHandler
	// 用函数嵌套的方式，实现了中间件的装饰模式 -> 问题？嵌套太长，没法批量设置
	core.Get("/user/login", framework.TimeoutHandler(UserLoginController, time.Second))
	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		// 动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

		// 嵌套路由
		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectNameController)
		}
	}
}
