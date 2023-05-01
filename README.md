# Pre-requisite

Run `git clone https://github.com/harshi93/go-movies-crud $HOME/go/src/`

Run `cd $HOME/go/src/go-movies-crud/`

Run `go get github.com/gorilla/mux`

# Compilation
`go build`

# Starting server

`./go-movies-crud`

# Accessing server
Server runs at localhost on port 8080

# Endpoints

#### Get all movies
`/movies`

#### Get movie by id
`/movie/{id}`

#### Add movie
`/addmovie`

#### Update movie
`/modmovie/{id}`

#### Delete movie
`/delmovies/{id}`