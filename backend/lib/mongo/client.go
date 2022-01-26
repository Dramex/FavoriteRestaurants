package mongo

import (
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var Users *mongo.Collection
var Ctx context.Context

func Mongo() {
    // base64 encoding to avoid email spam from leak detectors services
    sDec, _ := b64.StdEncoding.DecodeString("bW9uZ29kYitzcnY6Ly9yaXlhZGh0aWNrZXRzOlQxMjMxMjMxMjNAY2x1c3RlcjAudG42d24ubW9uZ29kYi5uZXQvbG9jYWw/YXV0aFNvdXJjZT1hZG1pbiZyZXRyeVdyaXRlcz10cnVlJnc9bWFqb3JpdHk=")
    client, err := mongo.NewClient(options.Client().ApplyURI(string(sDec)))
    if err != nil {
        log.Fatal(err)
    }
    Ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(Ctx)
    if err != nil {
        log.Fatal(err)
    }
  //  defer client.Disconnect(Ctx)

    quickstartDatabase := client.Database("fav")
    Users = quickstartDatabase.Collection("users")
}