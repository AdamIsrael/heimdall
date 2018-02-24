package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    // "github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
    "github.com/adamisrael/heimdall/config"
)

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
	// log.Printf("database uri is %s", configuration.Database.ConnectionUri)
	log.Printf("port for this application is %d", configuration.Server.Port)

    // viper.SetConfigType("yaml")
    //
    // viper.WatchConfig()
    // viper.OnConfigChange(func(e fsnotify.Event) {
    // 	fmt.Println("Config file changed:", e.Name)
    // })

    flag.Parse()
    fmt.Println("Verbose:", *verbose)

    // Start handling requests
    router := NewRouter()
    log.Printf("Listening on port %d", configuration.Server.Port)

    log.Fatal(http.ListenAndServe(
        fmt.Sprintf(":%d", configuration.Server.Port),
        router))

}
