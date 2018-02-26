package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    // "github.com/fsnotify/fsnotify"
    "github.com/adamisrael/heimdall/config"
    "github.com/spf13/viper"

    "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type ApiClient struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func main() {

    verbose := flag.Bool("verbose", false, "verbosity")

    // Setup configuration
    viper.SetConfigName("heimdall")
	viper.AddConfigPath(".")
	var configuration config.Configuration

    if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

    // viper.WatchConfig()
    // viper.OnConfigChange(func(e fsnotify.Event) {
    // 	fmt.Println("Config file changed:", e.Name)
    // })

    // Init database
    db, err = gorm.Open("sqlite3", "./heimdall.db")
    if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
    db.AutoMigrate(&ApiClient{})

	log.Printf("port for this application is %d", configuration.Server.Port)

    flag.Parse()
    fmt.Println("Verbose:", *verbose)

    // Start handling requests
    router := NewRouter()
    log.Printf("Listening on port %d", configuration.Server.Port)

    log.Fatal(http.ListenAndServe(
        fmt.Sprintf(":%d", configuration.Server.Port),
        router))

}
