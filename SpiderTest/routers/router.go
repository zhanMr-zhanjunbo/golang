package routers

import (
	"SpiderTest/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/movie_spider",&controllers.MovieSpiderController{},"*:GetSpider")
}
