package main

import (
    "fmt"
    "log"
    "net/http"
	"os"
    "github.com/gorilla/mux"
	"io/ioutil"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/post/{todoId}", TodoShow)

	port := os.Getenv("PORT")
    if port == "" {
	port = "8091"
    }	
    log.Fatal(http.ListenAndServe(":"+port, router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	dat, _ := ioutil.ReadFile("posts.txt");
    fmt.Fprintln(w, "Welcome!\n"+string(dat))
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
	f, err := os.OpenFile("posts.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(todoId); err != nil {
		panic(err)
	}
    fmt.Fprintln(w, "Todo show:", todoId)
}