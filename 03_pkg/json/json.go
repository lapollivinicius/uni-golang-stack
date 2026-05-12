package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Server ServerConfig `json:"server"`
}

type ServerConfig struct {
	Port int    `json:"port"`
	Host string `json:"host"`
}

func main() {
	
	// take json file
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	// close the read file
	defer file.Close()

	// decoder a content
	var data Config
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		panic(err)
	}

	fmt.Print(data.Server.Host)
	fmt.Println(data.Server.Port)

}