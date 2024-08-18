package routes

import (
	"github.com/fuatanshori/gin-gorm/controllers/book_controller"
	"github.com/fuatanshori/gin-gorm/controllers/user_controlller"
	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app
	route.GET("/user",user_controlller.GetAllUser)
	route.GET("/book",book_controller.GetAllBook)
}
