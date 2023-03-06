package main

import (
	"context"

	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/web"
	// 添加gin web服务starter
	_ "github.com/go-spring/starter-gin"
	// 添加echo web服务starter, 与gin web服务二选一即可
	//_ "github.com/go-spring/starter-echo"
	"log"
)

func init() {
	// 创建一个路由为/get-mapping的get请求资源, gin风格
	gs.GetMapping("/get-mapping", func(ctx web.Context) {
		ctx.JSON(web.SUCCESS.Data("get-mapping"))
	})
	// 创建一个路由为/get-binding的get请求资源, 对象风格
	gs.GetBinding("/get-binding", func(ctx context.Context, req *GetBindingReq) *GetBindingResp {
		return &GetBindingResp{Say: "hello " + req.Name}
	})
}

type GetBindingReq struct {
	Name string `form:"name" json:"name"`
}

type GetBindingResp struct {
	Say string `json:"say"`
}

func main() {
	log.Fatalln(gs.Run())
}
