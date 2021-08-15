package src

import (
	"encoding/json"
	"os"
)

func GetConfig(v string) string {
	f, err := os.Open("config/config.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(f)
	var token map[string]string
	err = decoder.Decode(&token)
	if err != nil {
		panic(err)
	}
	return token[v]
}
