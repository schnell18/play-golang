# Introduction

Simple golang CRUD server w/ [gorilla mux][1].

# Test case

## list all movies

    curl http://localhost:8000/movies/

## find a movie

    curl http://localhost:8000/movies/1

    curl http://localhost:8000/movies/2

## create movie

    curl -H "Content-Type: application/json" \
        http://localhost:8000/movies \
        -d'
        {
            "isbn": "44525235",
            "title": "Tron",
            "director": {
                "firstname": "Clinton",
                "lastname": "Eastwood"
            }
        }
        '
## update movie

    curl -XPUT -H "Content-Type: application/json" \
        http://localhost:8000/movies/8498081 \
        -d'
        {
            "isbn": "44525235",
            "title": "Tron 2021",
            "director": {
                "firstname": "Clinton",
                "lastname": "Eastwood"
            }
        }
        '

## delete movie

    curl -XDELETE http://localhost:8000/movies/1

[1]: https://github.com/gorilla/mux

[1]: https://github.com/gorilla/mux
