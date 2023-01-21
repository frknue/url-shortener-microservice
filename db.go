// db.go
package main

import (
    "fmt"
    "os"
    "log"
    "context"
    "time"
    "math/rand"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func createShortURL() string {
    rand.Seed(time.Now().UnixNano())
    var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, 5)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func InsertURL(url string) map[string]string {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
    uri := os.Getenv("MONGO_URI")
    if  uri == "" {
        log.Fatal("MONGO_URI is not set")
    }
    client , err := mongo.NewClient(options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal(err)
    }
    err = client.Connect(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    doc := map[string]string{
        "original_url": url,
        "short_url": createShortURL(),
    }

    collection := client.Database("url-shortener").Collection("urls")

    result, err := collection.InsertOne(context.Background(), doc)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result)
    return doc
}

func GetShortURL(url string) map[string]string {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
    uri := os.Getenv("MONGO_URI")
    if  uri == "" {
        log.Fatal("MONGO_URI is not set")
    }
    client , err := mongo.NewClient(options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal(err)
    }
    err = client.Connect(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    collection := client.Database("url-shortener").Collection("urls")

    var result map[string]string
    err = collection.FindOne(context.Background(), map[string]string{"short_url": url}).Decode(&result)
    if err != nil {
        log.Fatal(err)
    } 
    return result
}
