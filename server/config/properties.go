package config

import (
	"log"
	"math/rand"
	"os"
)

func GetallowAnonymousUsers() bool {
	return props.GetBool("allowAnonymousUsers", false)
}

func GetPrivateKey() string {
	return props.GetString("privateKey", "")
}

func GeneratePrivateKey() {
	//generate a private key using crypto/rand
	privateKey := make([]byte, 32)
	_, err := rand.Read(privateKey)
	if err != nil {
		log.Fatal("Error generating private key")
	}

	//append to data/server.properties file
	file, err := os.OpenFile("data/server.properties", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 600)
	if err != nil {
		log.Fatalln("Error while generating the private key: ", err)
	}
	_, err = file.WriteString("\nprivateKey=" + string(privateKey))
	if err != nil {
		log.Fatalln("Error while generating the private key: ", err)
	}

	_, _, err = props.Set("privateKey", string(privateKey))
	if err != nil {
		log.Fatalln("Failed to generate private key, try to restart: \n", err)
	}
}
