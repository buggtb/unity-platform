package routers

import (
	"unity-platform/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/install/:id", &controllers.InstallController{})
}
