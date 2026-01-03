package main

import (
    "context"
    "log"
    "net/http"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// Global variables
var (
    session *mongo.Client
    collection  *mongo.Collection
)

func init() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Connection to MongoDB
    client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal("Error al conectar a MongoDB:", err)
    }

    if err := client.Ping(ctx, nil); err != nil {
        log.Fatal("Ping a MongoDB fallido:", err)
    }

    // Store session and collection
    session = client
    collection = client.Database("MyDatabase").Collection("movies")

    log.Println("Conectado a MongoDB")
}

func main() {
    router := NewRouter()

    server := http.ListenAndServe(":8080", router)
    log.Fatal(server)

    // Cerrar MongoDB
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := session.Disconnect(ctx); err != nil {
        log.Fatal("Error al cerrar MongoDB:", err)
    }
    log.Println("Conexión a MongoDB cerrada")
}
