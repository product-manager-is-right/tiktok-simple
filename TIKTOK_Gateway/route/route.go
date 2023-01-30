package route

import (
	"TIKTOK_Gateway/configs"
	"TIKTOK_Gateway/resolver"
	"github.com/cloudwego/hertz/pkg/app/server"
)

const (
	GET  = "GET"
	POST = "POST"
)

// Register
/*
	根据微服务名，创建反向代理，并将config中同服务名的api注册到反向代理中。
*/
func Register(myConfig *configs.Config, h *server.Hertz) {
	// 创建反向代理
	proxyMap := resolver.CreateProxy(myConfig)

	// 将同一个服务的api注册到proxy中
	for _, route := range myConfig.Routes {
		switch route.Method {
		// 可以添加更多请求方式
		case GET:
			for _, api := range route.Apis {
				h.GET(api, proxyMap[route.ServiceName].ServeHTTP)
			}

		case POST:
			for _, api := range route.Apis {
				h.POST(api, proxyMap[route.ServiceName].ServeHTTP)
			}
		}
	}
}
