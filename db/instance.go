package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Instance struct {
	ColName string
	DBName  string

	db  *mongo.Database
	Col *mongo.Collection
}

var DBContext context.Context

func getDBContext() context.Context {
	if DBContext == nil {
		DBContext, _ = context.WithTimeout(context.Background(), 10*time.Second)
	}
	return DBContext
}
func (i *Instance) ApplyDatabase(db *mongo.Database) *Instance {
	i.db = db
	i.Col = db.Collection(i.ColName)
	i.DBName = db.Name()
	return i
}

func CreateUniversalDB(uri string, db string) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err.Error())
	}
	err = client.Connect(getDBContext())
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(db)
}
