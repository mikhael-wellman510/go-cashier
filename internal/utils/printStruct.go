package utils

import (
	"encoding/json"
	"log"
)

func PrintStruct(data any) {

	b, err := json.Marshal(data)

	if err != nil {
		log.Println("Error marshal")
		return
	}

	log.Println(string(b))

}
