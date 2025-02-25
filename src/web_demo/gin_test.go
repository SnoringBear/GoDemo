package web_demo

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"testing"
)

/**
Gin是一个用Go语言编写的轻量级Web框架，它凭借高性能、简洁易用的API、中间件支持和灵活的路由管理，
成为了许多开发者构建Web应用和微服务的首选工具。以下是对Gin框架的详细介绍：

一、主要特点
	高性能：Gin框架在内部使用了httprouter来实现路由功能，确保了快速的请求处理速度。据官方的性能测试数据，Gin可以处理每秒数十万次的请求。
	此外，Gin框架使用了Goroutine来提高并发性能，因此可以处理大量的并发请求。

	简洁易用的API：Gin提供了非常简洁且直观的API，开发者可以用极少的代码构建功能丰富的Web应用。

	中间件支持：Gin框架支持中间件的使用，中间件可以用来对请求做一些校验、过滤、鉴权的操作。通过编写和利用自定义中间件，可以进一步扩展Gin的功能，以满足特定的业务需求。

	路由管理：Gin框架支持多种路由方式，包括GET、POST、PUT、DELETE等，还支持参数的动态路由和路由组的嵌套。高效的路由匹配算法（基于树）使得Gin能够快速找到对应的处理器函数。


	gate -> web ?
*/

func TestGin01(t *testing.T) {
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()
	// 需要设置engine.SetTrustedProxies,否则[WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
	// 在 Gin 框架中，SetTrustedProxies 是一个设置函数，用于指定哪些代理可以信任并发送请求头 X-Forwarded-For 和 X-Real-IP。
	// 当使用 Gin 构建 API 服务器时，通常部署在反向代理（例如 Nginx）后面。客户端的请求通过反向代理转发到 Gin 服务器，
	// Gin 服务器通过 X-Forwarded-For 和 X-Real-IP 获取客户端的真实 IP 地址
	err := engine.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return
	}
	engine.GET("/test", ginHandler01)
	err = engine.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}

func ginHandler01(context *gin.Context) {
	remoteIP := context.RemoteIP()
	// context.RemoteIP() 无代理返回客户端IP,有代理返回代理IP
	log.Info().Msgf("remote ip: %s", remoteIP)
	// context.ClientIP() 无论是否有代理，都会返回客户端IP
	clientIP := context.ClientIP()
	log.Info().Msgf("client ip: %s", clientIP)
	context.JSON(200, map[string]interface{}{
		"code": 200,
		"msg":  "ok",
	})
}
