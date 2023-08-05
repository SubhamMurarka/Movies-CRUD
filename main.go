package main

import(
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

// creating Movie struct
type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

// creating director struct
type Director struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}


var movies []Movie


// getMovies func. returns all the movies data in json format

func getMovies(w http.ResponseWriter, r *http.Request){
	// changing content type to application/json
	w.Header().Set("Content-Type","application/json")
	// encoding movies into json format
	json.NewEncoder(w).Encode(movies)
	return 
}

// deletMovie func. deletes the movie with particualr id. 
func deleteMovie(w http.ResponseWriter, r *http.Request){
	// changing content type to application/json
	w.Header().Set("Content-Type","application/json")
	// collecting all the variables from url storing in map data structure.
	params := mux.Vars(r)
	for index , item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...) // unpacking the movies[index+1:] and appending into movies[:index] not including the movies[index]
			break
		}
	}
	// encoding movies into json format
	json.NewEncoder(w).Encode(movies)
	return
}

// getMovie func. returns the movie data with particular id in json format
func getMovie(w http.ResponseWriter , r *http.Request){
	// changing content type to application/json
	w.Header().Set("Content-Type","application/json")
	// collecting all the variables from url & storing in map data structure.
	params := mux.Vars(r)
	for _,item := range movies{
		if item.ID == params["id"]{
			// encoding movie with given id -> {item} into json format.
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// createMovie func. cretaes the new movie and add to database(movies) and returns the data in json format
func createMovie(w http.ResponseWriter , r *http.Request){
	// changing content type to application/json
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	// decoding the body of the response and storing in movie(by reference)
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// generating random id for the movie
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	// appending the newly created movie (newMovie) into movies dataset
	movies = append(movies,movie)
	// encoding newMovie into json format.
	json.NewEncoder(w).Encode(movie)
}

// updateMovie func. update the existing movie with new data
func updateMovie(w http.ResponseWriter , r *http.Request){
	// changing content type to application/json
	w.Header().Set("Content-Type","application/json")
	// collecting all the variables from url & storing in map data structure.
	params := mux.Vars(r)
	for index,movieValue := range movies{
		if movieValue.ID == params["id"]{
			// unpacking the movies[index+1:] and appending into movies[:index] not including the movies[index]
			movies = append(movies[:index],movies[index+1:]...)
			
			var newMovie Movie
			// decoding the body of the response and storing in movie(by reference)
			_ = json.NewDecoder(r.Body).Decode(&newMovie)
			newMovie.ID = params["id"]
			// appending the newly created movie (newMovie) into movies dataset
			movies = append(movies,newMovie)
			// encoding newMovie into json format.
			json.NewEncoder(w).Encode(newMovie)
		
		}
	}
	
}

func main(){
	// declaring a router
	r := mux.NewRouter()
	// making a small database 
	// database contains movie title, director name , id.
	movies  = append(movies ,Movie{ID:"1", Isbn:"438227", Title:"Movie One", Director: &Director{FirstName:"John",LastName:"Doe"}})
	movies  = append(movies ,Movie{ID:"2", Isbn:"45455", Title:"Movie Two", Director: &Director{FirstName:"Steve",LastName:"Smith"}})
	// different end points and their handling functions
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))

}