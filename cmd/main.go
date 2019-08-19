package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	proto "github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Descritpion", Content: "Hello World Article"},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Test Post Articles Endpoint")
	fmt.Fprint(w, "Endpoint Hit: Test Post Articles Endpoint")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Hello Camp Bowman")

	handleRequests()

	// Playing around with protoc
	kyle := &Person{
		Name: "Kyle",
		Age:  28,
		SocialFollowers: &SocialFollowers{
			Twitter: 500,
			Youtube: 5000,
		},
	}

	data, err := proto.Marshal(kyle)
	if err != nil {
		log.Fatal("Marshalling error: ", err)
	}

	fmt.Println(data)

	newKyle := &Person{}
	err = proto.Unmarshal(data, newKyle)
	if err != nil {
		log.Fatal("Unmarshalling error: ", err)
	}

	fmt.Println(newKyle.GetAge())
	fmt.Println(newKyle.GetName())
	fmt.Println(newKyle.SocialFollowers.GetTwitter())
	fmt.Println(newKyle.SocialFollowers.GetYoutube())
}
