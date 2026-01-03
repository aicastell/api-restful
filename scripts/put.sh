#! /bin/bash

curl -X PUT http://localhost:8080/pelicula/6952503aefee87baf16e76e4 \
    -H "Content-Type: application/json" \
    -d '{
        "name": "Harry el sucio",
        "year": 1971,
        "director": "Don Siegel"
    }'

