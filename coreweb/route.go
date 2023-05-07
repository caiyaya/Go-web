package main

import "coreweb/framework"

//func registerRouter(core *framework.Core) {
//	// 设置控制器
//	core.Get("foo", FooControllerHandler)
//}

// 注册路由
func registerRouter(core *framework.Core) {
	// http方法 + 静态路由匹配
	core.Get("user/login", UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		// 动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)
	}
}
