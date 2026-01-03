package main

import (
    "context"
    "fmt"
    "time"
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

func Index (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hola mundo desde mi servidor Go")
}

// Receives a Movie
func responseMovie(w http.ResponseWriter, status int, result Movie){
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(result)
}

// Receives an array of Movies
func responseMovies(w http.ResponseWriter, status int, results []Movie){
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(results)
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
    // Get ID from URL
    params := mux.Vars(r)
    movie_id := params["id"]

    // Convert string ID to ObjectID
    oid, err := primitive.ObjectIDFromHex(movie_id)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    // Create context
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Search document by _id
    var result Movie
    err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&result)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            http.Error(w, "Película no encontrada", http.StatusNotFound)
        } else {
            http.Error(w, "Error en la base de datos", http.StatusInternalServerError)
        }
        return
    }

    responseMovie(w, http.StatusOK, result)
}

func MovieList(w http.ResponseWriter, r *http.Request) {
   // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Exec query
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        http.Error(w, "Error al consultar MongoDB", http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx) // This is importat, free resources

   // Read all documents and store them into an slice
    var results []Movie
    if err = cursor.All(ctx, &results); err != nil {
        http.Error(w, "Error al procesar resultados", http.StatusInternalServerError)
        return
    }

    responseMovies(w, http.StatusOK, results)
}

// Define a decoder variable that will receive the data from the POST request.
// The request data, encoded in JSON, comes to us via r.Body.
// We convert it into an object we can work with.
func MovieAdd(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()

    var movie_data Movie
    err := json.NewDecoder(r.Body).Decode(&movie_data)

    if(err != nil) {
        http.Error(w, "JSON inválido", http.StatusBadRequest)
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err = collection.InsertOne(ctx, movie_data)
    if err != nil {
        http.Error(w, "Error interno", http.StatusInternalServerError)
        return
    }

    responseMovie(w, http.StatusOK, movie_data)
}

func MovieUpdate(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close() // Close connection

    // Get ID from URL
    params := mux.Vars(r)
    movie_id := params["id"]

    // Convert string to ObjectID
    oid, err := primitive.ObjectIDFromHex(movie_id)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    // Decode JSON from body
    var movie_data Movie
    if err := json.NewDecoder(r.Body).Decode(&movie_data); err != nil {
        http.Error(w, "JSON inválido", http.StatusBadRequest)
        return
    }

    // Create context
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Prepare fields to update (this is mongo syntax)
    update := bson.M{
        "$set": bson.M{
            "name":     movie_data.Name,
            "year":     movie_data.Year,
            "director": movie_data.Director,
        },
    }

    // Exec the update
    result, err := collection.UpdateOne(ctx, bson.M{"_id": oid}, update)
    if err != nil {
        http.Error(w, "Error al actualizar en la base de datos", http.StatusInternalServerError)
        return
    }

    // Verify if some document updated
    if result.MatchedCount == 0 {
        http.Error(w, "Película no encontrada", http.StatusNotFound)
        return
    }

    // Return the updated movie
    responseMovie(w, http.StatusOK, movie_data)
}

type Message struct {
    Status string `json:"status"`
    Message string `json:"message"`
}

func MovieRemove(w http.ResponseWriter, r *http.Request) {
    // Get the ID from URL
    params := mux.Vars(r)
    movie_id := params["id"]

    // Convert the string to ObjectID
    oid, err := primitive.ObjectIDFromHex(movie_id)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    // Create context
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Remove document by _id
    res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
    if err != nil {
        http.Error(w, "Error en la base de datos", http.StatusInternalServerError)
        return
    }

    if res.DeletedCount == 0 {
        http.Error(w, "Película no encontrada", http.StatusNotFound)
        return
    }

    // Prepare response
    result := Message{
        Status:  "success",
        Message: "La pelicula con ID " + movie_id + " ha sido borrada",
    }

    // Set headers and send answer
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(result)
}
