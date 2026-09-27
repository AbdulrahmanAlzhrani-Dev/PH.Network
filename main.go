package main

import (
	"fmt"
	"log"
	"net/http"
)

type AuthW struct {
	Name     string
	Password string
}

func RouterAPI() {
	mux := http.NewServeMux()
	auth := &AuthW{}

	mux.HandleFunc("/{$}", homePage)
	mux.HandleFunc("/create-user", CreateUser)
	mux.HandleFunc("POST /auth", auth.Auth)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintln(w, "WELCOME TO PRIVATE HYBRID NETWORK")
	fmt.Fprintln(w, name)
}

func (A *AuthW) Auth(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "<h1> THIS IS ENDPOINT IS NOT COMPLITED !</h1>")

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		jsonErr(w, 405, "use POST")
		return
	}
	id := r.URL.Query().Get("id")

	fmt.Fprint(w, id)

}

func jsonErr(w http.ResponseWriter, i int, s string) {
	panic("unimplemented")
}

func main() {

	fmt.Println("SERVER is run.... ; PORT:1000")
	RouterAPI()
}
