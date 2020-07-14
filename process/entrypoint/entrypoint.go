package entrypoint

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
	"usermanager/process/handler"
	"usermanager/process/schema"
)

//Generate id
func RandomSequence(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//Add user
func Add(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("begin")
	var user schema.User
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		resp := schema.Response{500, "payload error : " + err.Error()}
		rw.WriteHeader(500)
		json.NewEncoder(rw).Encode(resp)
		return
	}
	err = json.Unmarshal(b, &user) //payload obtained
	if err != nil {
		resp := schema.Response{500, "error : " + err.Error()}
		rw.WriteHeader(500)
		json.NewEncoder(rw).Encode(resp)
		return
	}
	user.Id = RandomSequence(16)
	err = handler.CreateHandler(user)
	if err != nil {
		resp := schema.Response{500, "error : " + err.Error()}
		rw.WriteHeader(500)
		json.NewEncoder(rw).Encode(resp)
		return
	}
	resp := schema.Response{200, "1 Row Inserted"}
	json.NewEncoder(rw).Encode(resp)
}

//List users
func List(rw http.ResponseWriter) {
	//fmt.Println("begin")
	users, err := handler.ListHandler()
	if err != nil {
		resp := schema.Response{500, "error : " + err.Error()}
		rw.WriteHeader(500)
		json.NewEncoder(rw).Encode(resp)
		return
	}
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(users)
}

//Read user by id
func Read(id string, rw http.ResponseWriter) {
	//fmt.Println("begin")
	users, err := handler.ReadHandler(id)
	if err != nil {
		resp := schema.Response{500, "error : " + err.Error()}
		rw.WriteHeader(500)
		json.NewEncoder(rw).Encode(resp)
		return
	}
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(users)
}

//Delete user by id
func Delete(id string, rw http.ResponseWriter) {
	err := handler.DeleteHandler(id)
	if err != nil {
		resp := schema.Response{500, "error : " + err.Error()}
		rw.WriteHeader(500)
		json.NewEncoder(rw).Encode(resp)
		return
	}
	resp := schema.Response{200, "id = " + id + " is deleted"}
	json.NewEncoder(rw).Encode(resp)
}

//Update user by id
func Update(id string, rw http.ResponseWriter, req *http.Request) {
	var user schema.User
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		resp := schema.Response{500, "payload error : " + err.Error()}
		rw.WriteHeader(500)
		json.NewEncoder(rw).Encode(resp)
		return
	}
	err = json.Unmarshal(b, &user) //payload obtained
	if err != nil {
		resp := schema.Response{500, "error : " + err.Error()}
		rw.WriteHeader(500)
		json.NewEncoder(rw).Encode(resp)
		return
	}
	user.Id = id
	users, err := handler.UpdateHandler(user)
	if err != nil {
		resp := schema.Response{500, "error : " + err.Error()}
		rw.WriteHeader(500)
		json.NewEncoder(rw).Encode(resp)
		return
	}
	json.NewEncoder(rw).Encode(users)
	fmt.Println(users)
}
