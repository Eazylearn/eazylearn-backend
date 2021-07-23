package main

import (
	"os"

	"github.com/Cumbercubie/api"
	"github.com/Cumbercubie/controller"
	"github.com/Cumbercubie/db"
	"github.com/Cumbercubie/model"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

// create collection instance for each model
func onDBConnected(c *mongo.Database) {
	model.InitTestModel(c)
	model.InitUserModel(c)
}
func main() {
	// initServer()
	server := api.InitServer()
	godotenv.Load(".env")
	DB_URI := os.Getenv("DB_URI")
	EzLearnDB := db.CreateUniversalDB(DB_URI, "carie")
	onDBConnected(EzLearnDB)
	//
	server.SetGroup("/test", controller.TestControllerGroup)
	server.SetGroup("/user", controller.UserControllerGroup)
	//
	server.Start(":8081")

}
