package directory

import (
	"fmt"
	"io/ioutil"
	"log"
	"usermanager/util/context"
	"usermanager/util/json"
)

//Reads neccesary configuration data from config file
func File(env string) {
	fmt.Println("begin")
	file, err := ioutil.ReadFile("config/" + env + ".json")
	if err != nil {
		log.Fatalf("Configuraiton '%s' file not found", "config/"+env+".json")
		return
	}
	data := json.Parse(file)
	context.Instance().Set("host", data.GetString("host"))
	context.Instance().Set("port", data.GetString("port"))
	context.Instance().Set("user", data.GetString("user"))
	context.Instance().Set("password", data.GetString("password"))
	context.Instance().Set("database", data.GetString("database"))
}
