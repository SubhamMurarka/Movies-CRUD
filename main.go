package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// creating Movie struct
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// creating director struct
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies = []Movie{
	{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}},
	{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{FirstName: "Steve", LastName: "Smith"}},
}

// getMovies func. returns all the movies data in json format

func getMovies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, movies)
}

// deletMovie func. deletes the movie with particualr id.
func deleteMovie(c *gin.Context) {
	id := c.Param("id")
	_, err := getMovieByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found"})
		return
	}
	_, err = removeMovieByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotImplemented, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, movies)
}

// getMovie func. returns the movie data with particular id in json format
func getMovie(c *gin.Context) {
	id := c.Param("id")
	movie, err := getMovieByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Movie Not Found"})
		return
	}
	c.IndentedJSON(http.StatusOK, movie)
}
func getMovieByID(id string) (*Movie, error) {
	for i, m := range movies {
		if m.ID == id {
			return &movies[i], nil
		}
	}
	return nil, errors.New("book not found")
}

// createMovie func. cretaes the new movie and add to database(movies) and returns the data in json format
func createMovie(c *gin.Context) {
	var newMovie Movie

	err := c.BindJSON(&newMovie)

	if err != nil {
		return
	}
	// generating random id for the movie
	newMovie.ID = strconv.Itoa(rand.Intn(10000000))
	// appending the newly created movie (newMovie) into movies dataset
	movies = append(movies, newMovie)
	c.IndentedJSON(http.StatusCreated, newMovie)
}

// updateMovie func. update the existing movie with new data
func updateMovie(c *gin.Context) {
	id := c.Param("id")
	_, err := getMovieByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found"})
		return
	}
	_, err = removeMovieByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotImplemented, gin.H{"message": "Movie not removed"})
		return
	}
	var newMovie Movie
	newMovie.ID = id
	err = c.BindJSON(&newMovie)
	if err != nil {
		return
	}
	movies = append(movies, newMovie)
	c.IndentedJSON(http.StatusOK, newMovie)
}

func removeMovieByID(id string) (*[]Movie, error) {
	for ind, movieValue := range movies {
		if movieValue.ID == id {
			movies = append(movies[:ind], movies[ind+1:]...)
			return &movies, nil
		}
	}
	return &movies, errors.New("not done")
}

func main() {
	// declaring a router
	r := gin.Default()
	// different end points and their handling functions
	r.GET("/movies", getMovies)
	r.GET("/movies/:id", getMovie)
	r.POST("/movies", createMovie)
	r.PUT("/movies/:id", updateMovie)
	r.DELETE("/movies/:id", deleteMovie)
	fmt.Printf("Starting server at port 8000\n")
	r.Run("localhost:8080")

}
