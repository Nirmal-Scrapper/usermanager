package schema

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
}
