package web

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//解析请求中带的参数
func GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	uid := params.ByName("uuid")
	fmt.Fprintf(w, "获取到用户ID:%s\n", uid)

	//获取请求方式
	mo := r.Method
	fmt.Fprintf(w, "获取到请求方式为:%s", mo)
}
