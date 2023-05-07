package framework

type IGroup interface {
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)
}

// Group struct 实现了IGroup
type Group struct {
	core   *Core
	prefix string
}

// 初始化 Group
func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		prefix: prefix,
	}
}

// 实现Get方法
func (g *Group) Get(uri string, handler ControllerHandler) {
	uri = g.prefix + uri
	g.core.Get(uri, handler)
}

// 实现PUT方法
func (g *Group) Put(uri string, handler ControllerHandler) {
	uri = g.prefix + uri
	g.core.Put(uri, handler)
}

// 实现POST方法
func (g *Group) Post(uri string, handler ControllerHandler) {
	uri = g.prefix + uri
	g.core.Post(uri, handler)
}

// 实现POST方法
func (g *Group) Delete(uri string, handler ControllerHandler) {
	uri = g.prefix + uri
	g.core.Delete(uri, handler)
}

//// 从 core中初始化这个Group
//// 这里 Group这个类型实现了IGroup接口，所以new的实例可以直接赋值给IGroup
//func (g *Group) Group(prefix string) IGroup {
//	return NewGroup(g.core, prefix)
//}
