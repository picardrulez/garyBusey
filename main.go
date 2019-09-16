package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"log"
	"net/http"
)

var VERSION = "0.1.0"
var DEFAULTAUTH = "nWo4lyfe"
var cryptKey = []byte("2531A1BBA018C535D8CDD231B95C8C0E")
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

var router = mux.NewRouter()

func main() {
	startup()
	router.HandleFunc("/", indexPageHandler)
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/logout", logoutHandler)
	router.HandleFunc("/admin", adminPageHandler)
	router.HandleFunc("/listUsers", listUsers)
	router.HandleFunc("/createUser", createUser)
	router.HandleFunc("/usercreation", usercreation).Methods("POST")
	router.HandleFunc("/changePassword", changePassword)
	router.HandleFunc("/passEntry", passEntry).Methods("POST")
	router.HandleFunc("/deleteUser", deleteUser)
	router.HandleFunc("/userdel", userdel).Methods("POST")
	router.HandleFunc("/addKey", addKey)
	router.HandleFunc("/createkey", createkey).Methods("POST")
	router.HandleFunc("/listKeys", listKeys)
	router.HandleFunc("/editKey", editKey)
	router.HandleFunc("/removeKey", removeKey)
	router.HandleFunc("/dropkey", dropkey).Methods("POST")

	http.Handle("/", router)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	log.Println("The Gary Busey Vault is now online")
	http.ListenAndServe(":8000", nil)
}
