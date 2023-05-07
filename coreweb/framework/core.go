package framework

import (
	"log"
	"net/http"
	"strings"
)

// 框架核心结构
type Core struct {
	//router map[string]ControllerHandler
	//router map[string]map[string]ControllerHandler
	router map[string]*Tree // all routers
}

// 初始化框架核心结构
func NewCore() *Core {
	//// 初始化router, 同时完成第一级hash的定义
	//getRouter := map[string]ControllerHandler{}
	//postRouter := map[string]ControllerHandler{}
	//putRouter := map[string]ControllerHandler{}
	//deleteRouter := map[string]ControllerHandler{}
	//
	//router := map[string]map[string]ControllerHandler{}
	//router["GET"] = getRouter
	//router["POST"] = postRouter
	//router["PUT"] = putRouter
	//router["DELETE"] = deleteRouter
	//return &Core{
	//	router: router,
	//}
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()
	return &Core{router: router}
}

//func (c *Core) Get(url string, handler ControllerHandler) {
//	c.router[url] = handler
//}

// 对应 Method = Get 哈希
//func (c *Core) Get(url string, handler ControllerHandler) {
//	upperUrl := strings.ToUpper(url)
//	c.router["GET"][upperUrl] = handler
//}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

// 对应 Method = Get Trie 树
func (c *Core) Get(url string, handler ControllerHandler) {
	if err := c.router["GET"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error :", err)
	}
}

// 对应 Method = Post Trie 树
func (c *Core) Post(url string, handler ControllerHandler) {
	if err := c.router["POST"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error :", err)
	}
}

// 对应 Method = Put Trie 树
func (c *Core) Put(url string, handler ControllerHandler) {
	if err := c.router["PUT"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error :", err)
	}
}

// 对应 Method = Deletex Trie 树
func (c *Core) Delete(url string, handler ControllerHandler) {
	if err := c.router["DELETE"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error :", err)
	}
}

//// 对应 Method = Get
//func (c *Core) Delete(url string, handler ControllerHandler) {
//	upperUrl := strings.ToUpper(url)
//	c.router["Delete"][upperUrl] = handler
//}

//// 对应 Method = Post
//func (c *Core) Post(url string, handler ControllerHandler) {
//	upperUrl := strings.ToUpper(url)
//	c.router["Post"][upperUrl] = handler
//}

//// 对应 Method = PUT
//func (c *Core) Put(url string, handler ControllerHandler) {
//	upperUrl := strings.ToUpper(url)
//	c.router["PUT"][upperUrl] = handler
//}

//// 路由匹配 哈希
//func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
//	// uri 和 method 全部转换为大写，保证大小写不敏感
//	uri := request.URL.Path
//	method := request.Method
//	upperMethod := strings.ToUpper(method)
//	upperUri := strings.ToUpper(uri)
//
//	// 第一层hash
//	if methodHandlers, ok := c.router[upperMethod]; ok {
//		// 第二层hash
//		if handler, ok := methodHandlers[upperUri]; ok {
//			return handler
//		}
//	}
//	return nil
//}

// 路由匹配 -> trie树
func (c *Core) FindRouterByRequest(request *http.Request) ControllerHandler {
	//uri 和 method 全部转换为大写，保证大小写不敏感
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)
	// 查找第一层map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.FindHandler(uri)
	}
	return nil
}

// 框架核心结构实现Handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("core.serevrHttp start.")
	// 封装自定义context
	ctx := NewContext(request, response)

	// 简单的路由选择，这里先写死
	//router := c.router["foo"]
	//if router == nil {
	//	return
	//}
	//log.Println("core.router success.")
	//router(ctx)

	// 寻找路由
	router := c.FindRouterByRequest(request)
	if router == nil {
		ctx.Json(404, "router not found")
		return
	}
	// 调用路由函数
	if err := router(ctx); err != nil {
		ctx.Json(500, "inner error")
		return
	}
}
