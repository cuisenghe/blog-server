package main

import (
	"blog-server/configs"
	"blog-server/internal/router"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	// 加载配置
	configs.InitConfig()
	// 初始化路由
	r := router.InitRouter()

	// run
	r.Run(":8080")
}
