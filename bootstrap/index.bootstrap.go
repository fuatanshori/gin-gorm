package bootstrap

import (
	"net/http"

	"github.com/fuatanshori/gin-gorm/configs/app_config"
	"github.com/fuatanshori/gin-gorm/db"
	"github.com/fuatanshori/gin-gorm/routes"
	"github.com/gin-gonic/gin"
)

func BootstrapApp() {
	db.ConnectDatabase()
	app := gin.Default()
	routes.InitRoute(app)
	err := http.ListenAndServe(app_config.APP_ADDRESS+":"+app_config.APP_PORT, app)
	if err != nil {
		panic(err)
	}
}
