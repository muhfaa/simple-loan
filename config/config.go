package env

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

// struct config like in config json
type AppConfig struct {
	BackendPort int `json:"backend_port"`
	MongoDB     struct {
		DBDriver   string `json:"dbdriver"`
		DBName     string `json:"dbname"`
		DBAddress  string `json:"dbaddress"`
		DBPort     string `json:"dbport"`
		DBUsername string `json:"dbusername"`
		DBPassword string `json:"dbpassword"`
	} `json:"mongodb"`
	Mysql struct {
		Host              string `json:"host"`
		Port              string `json:"port"`
		User              string `json:"user"`
		Password          string `json:"password"`
		Name              string `json:"name"`
		MaxIdleConnection int    `json:"max_idle_con"`
		MaxOpenConnection int    `json:"max_open_con"`
	} `json:"mysql"`

	MongoDatabase *mongo.Database
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

// GetConfig Initiatilize config in singleton way
func GetConfig() *AppConfig {
	if appConfig != nil {
		return appConfig
	}

	lock.Lock()
	defer lock.Unlock()

	// re-check after locking
	if appConfig != nil {
		return appConfig
	}

	appConfig = initConfig()

	return appConfig
}

func initConfig() *AppConfig {
	var configFile = "config.json"
	var finalConfig AppConfig

	jsonByte, err := ioutil.ReadFile(fmt.Sprintf("config/%s", configFile))
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonByte, &finalConfig)
	if err != nil {
		panic(err)
	}

	return &finalConfig

}
