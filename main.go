package main

import (
	"log"

	"mdsybd/models"

	_ "github.com/mattn/go-oci8"
)

func main() {
	client, err := models.NewClient()
	if err != nil {
		log.Fatalln("MainL create client err:", err)
	}
	defer client.Close()

}
