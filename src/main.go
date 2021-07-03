package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Leng-Kai/bow-code-API-server/db"
	"github.com/Leng-Kai/bow-code-API-server/routes"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
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
	r := routes.NewRouter()
	http.Handle("/", r)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},   // All origins
		AllowedMethods: []string{"GET"}, // Allowing only get, just an example
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), c.Handler(r)))
}
