# Go RESTful API Example with MongoDB

This repository demonstrates how to build a simple RESTful API in Go using MongoDB as the backend database and the Gorilla/Mux router.

It is designed as a minimal, educational example, focused on clarity, not production architecture. So developers can understand the core concepts of routing, request handling, and database integration in Go.

# Requirements

- Go: 1.24.5
- Docker (to run MongoDB)

# MongoDB setup

This API uses a MongoDB instance. The easiest way to run it is with Docker:

```
$ mkdir -p mongo-data
$ docker run -d \
    --name mongodb \
    -p 27017:27017 \
    -v ~/mongo-data:/data/db \
    mongo:7.0
```

The application connects to the database named *MyDatabase* and uses the collection *movies*. MongoDB creates both automatically on the first write. So, no manual setup is needed.

# Running the API

Clone the repository

```
$ git clone https://github.com/aicastell/api-restful.git
```

And run the server:

```
$ cd api-restful
$ go run .
```

The API will start on http://localhost:8080

# API endpoints

Method Endpoint Description
GET / Welcome message
GET /peliculas List all movies
GET /pelicula/{id} Get a movie by ID
POST /pelicula Create a new movie
PUT /pelicula/{id} Update an existing movie
DELETE /pelicula/{id} Delete a movie by ID

# Source files

- movie.go: Defines the data model (Movie)
- actions.go: Implements HTTP handlers
- routes.go: Registers the routes in the Gorilla/Mux router
- main.go: Connection to MongoDB and start HTTP server

# Helper scripts

Add a new movie:

```
./scripts/post.sh
```

Update a movie (replace <ID> in the script)

```
./scripts/put.sh
```

Delete a movie (replace <ID> in the script)

```
./scripts/delete.sh
```

# Dependencies

Managed via Go modules. Defined in go.mod:

- github.com/gorilla/mux – HTTP router
- go.mongodb.org/mongo-driver – MongoDB driver

No manual installation is required. Comand go run . will download them automatically.

# License

This project is licensed under the GPL Version 2

