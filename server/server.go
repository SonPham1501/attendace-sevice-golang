package server

import (
	"com.son.server.attendaceservice/configs"
	"com.son.server.attendaceservice/routers"
)

func Init() {
	config := configs.GetConfig()
	r := routers.NewRouter()
	r.Run(":" + config.GetString("server.port"))
}
