// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"fox/controllers/admin"
	"fox/controllers"
)

func init() {
	beego.Router("/", &controllers.BlogController{},"get:GetAll")
	beego.Router("/:id", &controllers.BlogController{},"get:Get")
	beego.Router("/page/:page", &controllers.BlogController{},"get:GetAll")
	//beego.Router("/admin/login", &admin.LoginController{})
	//beego.Router("/admin/index", &admin.IndexController{})
	ns := beego.NewNamespace("/admin",
		beego.NSRouter("/login", &admin.LoginController{}),
		beego.NSRouter("/logout", &admin.LogoutController{}),
		beego.NSRouter("/index", &admin.IndexController{}),
		beego.NSRouter("/index_v2", &admin.IndexV2Controller{}),
		beego.NSRouter("/my_password", &admin.MyPasswordController{}),
		//blog
		beego.NSRouter("/blogs", &admin.BlogController{},"get:List"),
		beego.NSRouter("/blog", &admin.BlogController{}),
		beego.NSRouter("/blog/:id", &admin.BlogController{}),//, "get:Get"
		beego.NSRouter("/blog/:id", &admin.BlogController{}),//, "put:Put"
		beego.NSRouter("/blog/edit/:id", &admin.BlogController{}, "get:Edit"),
		beego.NSRouter("/blog/add", &admin.BlogController{}, "get:Add"),
	)
	beego.AddNamespace(ns)
	//ns := beego.NewNamespace("/v1",
	//
	//	beego.NSNamespace("/admin",
	//		beego.NSInclude(
	//			&controllers.AdminController{},
	//		),
	//	),
	//)
	//beego.AddNamespace(ns)
}
