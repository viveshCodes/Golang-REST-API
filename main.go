package main

/*_______Import Section_______________
______________________________________*/
import (
	//	"encoding/json"
	"encoding/json"
	"log"
	"net/http"

	//	"math/rand"
	//	"strconv"
	"github.com/gorilla/mux" // To import this ,we should install it using ``` go get -u gthub.com/gorilla/mux ```
)


/*__________Models_______________
__________________________________*/

// Book Struct (Model)
type Book struct{
	ID      string  `json:"id"`
	Isbn    string  `json:"isbn"`
	Title   string  `json:"title"`
	Author  *Author `json:"author"`
}
// Author Struct
type Author struct{
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

/*_________________Slice______________________
______________________________________________*/
//Init books var as a slice Book struct
var books []Book


/*____________Router Functions________________
_______________________________________________*/

// getBooks
func getBooks(w http.ResponseWriter ,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)

}
//getBook
func getBook(w http.ResponseWriter ,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)  // Get params . In this context we'll get id

	// Loop through all books and find the correct id
	for _,item :=range books{     // range is used to loop through map,slice ,or any data structure       
		if item.ID == params ["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	} 
	json.NewEncoder(w).Encode(&Book{})
	
}
//createBook
func createBook(w http.ResponseWriter ,r *http.Request){
	
}
//updateBook
func updateBook(w http.ResponseWriter ,r *http.Request){
	
}
//deleteBook
func deleteBook(w http.ResponseWriter ,r *http.Request){
	
}


/*___________Main Function_________________
___________________________________________*/
func main(){
 //Init Router
 router :=mux.NewRouter()

 //Mock Data -@todo-implement DB
 books = append(books,Book{ID:"1",Isbn:"7781",Title:"Book One",Author:&Author{Firstname : "Vivesh",Lastname:"Yadav"}})
 books = append(books,Book{ID:"2",Isbn:"1981",Title:"Book Two",Author:&Author{Firstname : "vivesh",Lastname:"Codes"}})

 // Create Route Handlers / These route handlers will establish Endpoints for our API
 // HandleFunc("routes,function")
 router.HandleFunc("/api/books",getBooks).Methods("GET")
 router.HandleFunc("/api/books/{id}",getBook).Methods("GET")
 router.HandleFunc("/api/books",createBook).Methods("POST")
 router.HandleFunc("/api/books/{id}",updateBook).Methods("PUT")
 router.HandleFunc("/api/books/{id}",deleteBook).Methods("DELETE")

 // To run server
 log.Fatal(http.ListenAndServe(":8000",router)) // We've used log.Fatal() to catch error

}