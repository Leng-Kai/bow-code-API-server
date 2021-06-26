package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/J-HowHuang/bow-code/db"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := mux.NewRouter()
	db_url := os.Getenv("DB_URL")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db_url))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	if err == nil {
		err := client.Ping(ctx, nil)
		if err == nil {
			log.Println("DB connected at " + db_url)
		} else {
			log.Println("Failed to connect DB at " + db_url)
		}
	} else {
		log.Fatal(err)
	}
	db.InitDB(client)

	http.ListenAndServe(":8080", r)
}
