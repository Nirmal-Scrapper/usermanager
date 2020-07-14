package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

//User data schema
type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
	City  string `json:"city"`
}

// Response for api calls with error or mesage
type Response struct {
	Code     int
	Response string
	ID       string
	Error    string
}

func Tests(t *testing.T) {
	var user User
	var id string
	data, err := ioutil.ReadFile("newuser.json")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	err = json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	url := "http://localhost:8015/api/user"
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	res, _ := client.Do(req)
	fmt.Println(url)
	if res.StatusCode != 200 {
		t.Error("failed")
	}
	respond, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var response Response
	err = json.Unmarshal(respond, &response) //payload obtained
	if err != nil {
		fmt.Println(err)
	}
	if user.Id == "" {
		t.Error("failed")
	}

	//200 testcase
	id = user.Id
	url = "http://localhost:8015/api/user"
	client = &http.Client{}
	req, _ = http.NewRequest("GET", url, nil)
	res, _ = client.Do(req)
	fmt.Println(url)
	if res.StatusCode != 200 {
		t.Error("failed")
	}
	url = "http://localhost:8015/api/user/" + id
	client = &http.Client{}
	req, _ = http.NewRequest("GET", url, nil)
	res, _ = client.Do(req)
	fmt.Println(url)
	if res.StatusCode != 200 {
		t.Error("failed")
	}
	url = "http://localhost:8015/api/user/" + id
	client = &http.Client{}
	req, _ = http.NewRequest("PUT", url, bytes.NewBuffer(data))
	res, _ = client.Do(req)
	fmt.Println(url)
	if res.StatusCode != 200 {
		t.Error("failed")
	}

	url = "http://localhost:8015/api/user/" + id
	client = &http.Client{}
	req, _ = http.NewRequest("DELETE", url, nil)
	res, _ = client.Do(req)
	fmt.Println(url)
	if res.StatusCode != 200 {
		t.Error("failed")
	}

	//error test cases : sucessive call is failed
	url = "http://localhost:8015/api/user/1"
	client = &http.Client{}
	req, _ = http.NewRequest("GET", url, nil)
	res, _ = client.Do(req)
	fmt.Println(url)
	if res.StatusCode == 200 {
		t.Error("failed")
	}

	url = "http://localhost:8015/api/user/76"
	client = &http.Client{}
	req, _ = http.NewRequest("DELETE", url, nil)
	res, _ = client.Do(req)
	fmt.Println(url)
	if res.StatusCode == 200 {
		t.Error("failed")
	}
}
