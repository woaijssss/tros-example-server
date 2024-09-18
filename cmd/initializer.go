package cmd

import (
	"github.com/woaijssss/tros"
	trosConf "github.com/woaijssss/tros/conf"
	"github.com/woaijssss/tros/server/grpc"
	"github.com/woaijssss/tros/server/http"
	http2 "github.com/woaijssss/tros/server/middleware/http"
	"runtime"
	"tros_example_server/internal/api"
)

func GlobalInit() {
	//conf.InitConfig()	// 初始化读取自定义配置
	http2.Start(trosConf.GetAppName(), trosConf.GetMonitorPort())
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// HttpServerInit 同时支持gin版 http server初始化
//func HttpServerInit() {
//	e := http.NewEngine(
//		http2.Recovery(),
//		http2.Cors(),
//		http2.Monitor(),
//		http2.HeartCheck(),
//	)
//
//	http.AddMiddleWares(e,
//		gin.Logger(),
//		sentrygin.New(sentrygin.Options{Repanic: true}),
//	)
//	//e.SetTrustedProxies(nil)
//
//	// 初始化路由器
//	//e = router.InitRouter(e)
//
//	e.Run(fmt.Sprintf(":%d", trosConf.GetHttpPort()))
//}

// GrpcGatewayServerInit grpc-gateway版 http server初始化
func GrpcGatewayServerInit() {
	app := tros.New(
		tros.Servers(grpc.DefaultServer(), http.DefaultServer()),
		tros.WithInitializers([]tros.Initializer{
			api.ExampleController, // 每一个controller层在这里初始化
		}...),
	)

	app.Run()
}
