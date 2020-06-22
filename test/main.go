package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/dealancer/validate.v2"
)

type Config struct {
	Name       string   `yaml:"name"    json:"name"     validate:"empty=false"`
}

func main() {
	var myConfig Config

	var rawConfig = []byte(`{}`)  // Empty, Nothing getting passed.

	err := json.Unmarshal(rawConfig, &myConfig)
	if err != nil {
		panic(err)
	}
	err = validate.Validate(myConfig)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Name = %s\n", myConfig.Name)
}