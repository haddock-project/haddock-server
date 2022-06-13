package config

import (
	"log"
	"math/rand"
)

func GetallowAnonymousUsers() bool {
	return props.GetBool("allowAnonymousUsers", false)
}

func GetDebugMode() bool {
	return props.GetBool("debugMode", false)
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

	//update the property
	_, _, err = props.Set("privateKey", string(privateKey))
	if err != nil {
		log.Fatalln("Failed to generate private key, try to restart: \n", err)
	}

	//save the new property
	if err := Save(props); err != nil {
		log.Fatalln("Failed to save private key, try to restart: \n", err)
	}
}
