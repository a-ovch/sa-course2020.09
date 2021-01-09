package main

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

const appName = "user"

type Config struct {
	Port int `envconfig:"port"`
}

func main() {
	var c Config
	err := envconfig.Process(appName, &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Port: %v", c.Port)
}
