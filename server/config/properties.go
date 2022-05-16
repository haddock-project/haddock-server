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
	//generate random characters and set it as private key
	var privateKey string
	for i := 0; i < 32; i++ {
		privateKey += string(rune(65 + rand.Intn(25)))
	}
	//append to data/server.properties file
	file, err := os.OpenFile("data/server.properties", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 600)
	if err != nil {
		log.Fatalln("Error while generating the private key: ", err)
	}
	_, err = file.WriteString("\nprivateKey=" + privateKey)
	if err != nil {
		log.Fatalln("Error while generating the private key: ", err)
	}

	_, _, err = props.Set("privateKey", privateKey)
	if err != nil {
		log.Fatalln("Failed to generate private key, try to restart: \n", err)
	}
}
