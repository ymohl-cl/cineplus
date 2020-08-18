# CINEPLUS

CINEPLUS is a recruit test in Go.

``` txt
Go Back-end Assignment: Movie List

Studio Ghibli is a Japanese movie company. They offer a REST API where one can query information about movies and people (characters).

The task is to write a Go application which serves a page on localhost:8000/movies/. This page should contain a plain list of all movies from the Ghibli API. For each movie the people that appear in it should be listed.

Do not use the people field on the /films endpoint, since it’s broken. Instead, there is a list field called films on the /people endpoint which you can use to obtain the relationship between movies and the people appearing in them.

You don’t have to worry about the styling of that page.

Since accessing the API is a time-intensive operation, it should not happen on every page load. But on the other hand, movie fans are a very anxious crowd when it comes to new releases, so make sure that the information on the page is not older than 1 minute when the page is loaded.

The code should be submitted in a clean and refactored state.

Don’t forget to test your code. Your tests don’t have to be complete, but you should describe how you would extend them if you had the time.

If you have to skip some important work due to time limitations, feel free to add a short description of what you would improve and how if you had the time for it.
```

## Table of Contents

- [CINEPLUS](#cineplus)
  - [Table of Contents](#table-of-contents)
  - [Requirements](#requirements)
  - [Build usage](#build-usage)
    - [From Makefile](#from-makefile)
    - [From docker-compose](#from-docker-compose)
  - [API Response](#api-response)
  - [Github actions](#github-actions)
  - [HTTP Framework](#http-framework)
  - [Versionning](#versionning)
  - [Contact](#contact)

## Requirements

- GNU Make >= 3.81
- Golang >= 1.15
- Docker-Compose >= 1.25.4
- Docker >= 19.03.5

## Build usage

To run the cineplus app, you should set the following variables:

``` bash
export CINEPLUS_URL=https://ghibliapi.herokuapp.com
export CINEPLUS_REFRESHTIME=60000 # millisecond
export CINEPLUS_PORT=8000
```

***The default environments variables are set only with docker-compose solution***

### From Makefile

``` bash
# Download tools (golint, gomock ...)
make tools

# Download dependencies (go modules)
make install

# Linter
make lint

# Units tests
make test

# Build
make build

# Execute
./bin/app
```

### From docker-compose

``` bash
docker-compose up --remove-orphans --build
```

You can update the environments values from the docker-compose.yml file.

## API Response

``` json
{
  "movies":[
    {
      "identifier":"movie_identifier",
      "title":"movie_title",
      "description":"movie_description",
      "director":"movie_director",
      "producer":"movie_producer",
      "release_date":"movie_release_date",
      "score":"movie_score",
      "peopleIds":["people_identifier_1", "people_identifier_2"]
    }
  ],
  "peoples":[
    {
      "id":"people_identifier",
      "name":"people_name",
      "gender":"people_gender",
      "age":"people_age",
    }
  ]
}
```

## Github actions

A github action is executed at each push and pull request to pass the linter, the tests and the build process.

## HTTP Framework

I choice [Echo](github.com/labstack/echo) from labstack because it is a easy, complete and light framework. Of course, i could have used [gorilla mux](github.com/gorilla/mux) or other.

## Versionning

No versionning.

## Contact

mohl.clauzade@gmail.com
