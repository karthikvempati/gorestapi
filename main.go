package main

import ( 
	"fmt"
	"log"
	"net/http" 
	"encoding/json"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article

func getAllArticles(w http.ResponseWriter, r *http.Request){
	articles := Articles{
		Article{Title:"Karthik's article", Desc:"My first artile", Content:"This is a sample article to test the rest api"},
	}
	
	fmt.Println("Endpoint Hit: All articles endpoint")
	json.NewEncoder(w).Encode(articles)
}


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Endpoint Hit")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", getAllArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}