package routers

import (
	"ehang.io/cnt/web/controllers"
	"github.com/astaxie/beego"
)

func Init() {
	// 自定义错误处理
	beego.ErrorHandler("404", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusNotFound)
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		// 你可以渲染一个自定义的404页面或返回错误信息
		rw.Write([]byte("<html><body><h1>404 页面未找到</h1><p>抱歉，您访问的页面不存在。</p></body></html>"))
	})

	web_base_url := beego.AppConfig.String("web_base_url")
	if len(web_base_url) > 0 {
		ns := beego.NewNamespace(web_base_url,
			beego.NSRouter("/", &controllers.IndexController{}, "*:Index"),
			beego.NSAutoRouter(&controllers.IndexController{}),
			beego.NSAutoRouter(&controllers.LoginController{}),
			beego.NSAutoRouter(&controllers.ClientController{}),
			beego.NSAutoRouter(&controllers.AuthController{}),
		)
		beego.AddNamespace(ns)
	} else {
		beego.Router("/", &controllers.IndexController{}, "*:Index")
		beego.AutoRouter(&controllers.IndexController{})
		beego.AutoRouter(&controllers.LoginController{})
		beego.AutoRouter(&controllers.ClientController{})
		beego.AutoRouter(&controllers.AuthController{})
	}
}