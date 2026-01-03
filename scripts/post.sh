#! /bin/bash

curl -X POST http://localhost:8080/pelicula \
  -H "Content-Type: application/json" \
  -d '{"name":"Terminator","year":1984,"director":"James Cameron"}'

curl -X POST http://localhost:8080/pelicula \
  -H "Content-Type: application/json" \
  -d '{"name":"El nombre de la rosa","year":1986,"director":"Jean-Jacques Annaud"}'

curl -X POST http://localhost:8080/pelicula \
  -H "Content-Type: application/json" \
  -d '{"name":"Gladiator","year":2000,"director":"Ridley Scott"}'

curl -X POST http://localhost:8080/pelicula \
  -H "Content-Type: application/json" \
  -d '{"name":"Idiocracia","year":2006,"director":"Mike Judge"}'
