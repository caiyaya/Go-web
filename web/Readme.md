# 1.1 Http基础:实现了路由映射表，提供了用户注册静态路由的方法，包装了启动服务的函数
核心点：
    理解net/http库提供的Handle接口，它只包含方法ServeHttp，对于启动服务函数ListenAndServer(addr string, h Handler)而言，
只要传入了任何实现了ServerHttp方法的实例，所有的http请求，都可以从而被拦截处理，即统一控制入口，这也是自建框架的第一步。

# 1.2 上下文Context：
核心点：
    1、引入context的必要性，对于*http.Request和http.ResponseWriter 有很多重复的消息无需业务代码解析，框架应该提供相应的解析方式。
    2、在引入中间件后产生的信息应该统一处理，且对业务应该是无感知的，即可以封装好达到随去随用，且随请求的开始而开始，结束而结束，

# 1.3 前缀树路由Router:
核心点： 对于动态路由的理解 对于前缀树的理解 具体场景的应用？ 
     map的形式只能处理静态路由，无法对动态路由进行处理，因此引入另一种数据机构Trie树，实现对:name 和 *filepath两种路由方式的解析

# 1.4 分组控制Group:
核心点：实现路由的分组控制 Route Group Control, 以相同前缀进行分组

