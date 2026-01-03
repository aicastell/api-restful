# Introduccion

Este código muestra como implementar una API RESTFul en Golang usando MongoDB como backend.

# Golang

El router de la API RESTFul esta implementado en GoLang usando la librería Gorilla/MUX. Es un ejemplo muy básico centrado en la parte esencial del router y en los metodos del API.


# MongoDB

Como base de datos del backend voy a usar MongoDB. La instalación de MongoDB es muy sencilla, basta con instalar docker y ejecutar estos comandos:

    $ mkdir -p mongo-data

Ejecutar el contenedor:

    $ docker run -d \
        --name mongodb \
        -p 27017:27017 \
        -v ~/mongo-data:/data/db \
        mongo:7.0


El ejemplo hace uso de una base de datos llamada MyDatabase, que puedes crear siguiendo estos pasos:




Puedes conectar a la base de datos haciendo esto:

TODO.






El código define una estructura Movie, y crea varios metodos para:

        Route {
            "MovieList",
            "GET",
            "/peliculas",
            MovieList,
        },
        Route {
            "MovieShow",
            "GET",
            "/pelicula/{id}",
            MovieShow,
        },
        Route {
            "MovieAdd",
            "POST",
            "/pelicula",
            MovieAdd,
        },
        Route {
            "MovieUpdate",
            "PUT",
            "/pelicula/{id}",
            MovieUpdate,
        },
        Route {
            "MovieRemove",
            "DELETE",
            "/pelicula/{id}",
            MovieRemove,
        },

El código puede ser usado para investigar como se implementa dicha API RESTFUL en GoLang.

Pasos para compilar el proyecto:

go get github.com/gorilla/mux
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/bson


