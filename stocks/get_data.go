package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	s "strings"

	"github.com/gorilla/mux"
)

var strToken = "##token##"
var url = "https://www.quandl.com/api/v3/datasets/WIKI/" + strToken + "/data.json?api_key=keZt48nPpZ8cDeQZ8p2w"

type Tracks struct {
	Toptracks []Toptracks_info
}

type Toptracks_info struct {
	Track []Track_info
	Attr  []Attr_info
}

type Track_info struct {
	Name       string
	Duration   string
	Listeners  string
	Mbid       string
	Url        string
	Streamable []Streamable_info
	Artist     []Artist_info
	Attr       []Track_attr_info
}

type Attr_info struct {
	Country    string
	Page       string
	PerPage    string
	TotalPages string
	Total      string
}

type Streamable_info struct {
	Text      string
	Fulltrack string
}

type Artist_info struct {
	Name string
	Mbid string
	Url  string
}

type Track_attr_info struct {
	Rank string
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// json data
	url += "&limit=1" // limit data for testing
	res, err := http.Get(s.Replace(url, strToken, params["id"], -1))
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	var data interface{} // TopTracks
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Results: %v\n", data)
	json.NewEncoder(w).Encode(&data)
}

func main() {
	fmt.Printf("init succesfully. Come on!!!")
	router := mux.NewRouter()
	router.HandleFunc("/stock/{id}", GetStock).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
