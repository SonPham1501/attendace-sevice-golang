package server

import (
	"context"
	"log"

	"com.son.server.attendaceservice/configs"
	"com.son.server.attendaceservice/db/mongodb"
	"com.son.server.attendaceservice/routers"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Init() {
	config := configs.GetConfig()

	// Create database connection
	createDBConnection(config)

	// Init routers server
	r := routers.NewRouter()
	r.Run(":" + config.GetString("server.port"))
}

func createDBConnection(config *viper.Viper) {
	path := config.GetString("db.path")
	mongodb.NewMongoDB(path)

	err := mongodb.GetMongoDBClient().Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}
}
