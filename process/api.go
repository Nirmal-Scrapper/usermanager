package process

import (
	"fmt"
	"net/http"
	"usermanager/process/entrypoint"

	"github.com/gorilla/mux"
)

//List the users
func List(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("List")
	entrypoint.List(rw)
}

//Read users based on id
func ReadUser(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("begin")
	vars := mux.Vars(req)
	id := vars["id"]
	entrypoint.Read(id, rw)
}

//Create new user
func CreateUser(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("begin")
	entrypoint.Add(rw, req)
}

//Update user by Id
func UpdateUser(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("begin")
	vars := mux.Vars(req)
	id := vars["id"]
	entrypoint.Update(id, rw, req)
}

//Delete user by id
func DeleteUser(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("begin")
	vars := mux.Vars(req)
	id := vars["id"]
	entrypoint.Delete(id, rw)
}
