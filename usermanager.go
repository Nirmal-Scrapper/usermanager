package main

import (
	"fmt"
	"net/http"
	"os"
	"usermanager/directory"
	"usermanager/process"
	"usermanager/util/db"

	"github.com/gorilla/mux"
)

var env string

func main() {
	fmt.Println("begin")
	port := ":8015"
	if len(os.Args) > 1 {
		env = os.Args[1]
	} else {
		fmt.Println("no environment mentioned so local is chosen")
		env = "local"
	}
	directory.File(env)
	db.Connect()
	r := mux.NewRouter()
	r.HandleFunc("/api/status", GetStatus).Methods("GET")
	r.HandleFunc("/api/user", process.List).Methods("GET")
	r.HandleFunc("/api/user", process.CreateUser).Methods("POST")
	r.HandleFunc("/api/user/{id}", process.ReadUser).Methods("GET")
	r.HandleFunc("/api/user/{id}", process.DeleteUser).Methods("DELETE")
	r.HandleFunc("/api/user/{id}", process.UpdateUser).Methods("PUT")
	fmt.Println("running at ", port)
	http.ListenAndServe(port, r)
}

//Get the status of the service
func GetStatus(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(200)
	rw.Write([]byte("running"))
}
