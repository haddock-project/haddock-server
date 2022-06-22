package config

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
	"os"
	"strings"
	"time"
)

func GetallowAnonymousUsers() bool {
	return props.GetBool("allowAnonymousUsers", false)
}

func GetDebugMode() bool {
	return props.GetBool("debugMode", false)
}

func GetDockerSocketPath() string {
	return props.GetString("dockerSocketPath", "unix:///var/run/docker.sock")
}

func GetHost() string {
	return props.GetString("host", ":8080")
}

func GetTokenExpiration() time.Time {
	prop := props.GetString("tokenExpiration", "2h")

	//disable token expiration
	if strings.HasPrefix(prop, "0") {
		return time.Time{}
	}

	duration, err := time.ParseDuration(prop)
	if err != nil {
		//Saving the default setting
		log.Println("[WARNING] Invalid token expiration duration. Using default value of 2 hours.")
		props.Set("tokenExpiration", "2h")
		Save(props)

		//Return the default value
		return time.Now().Add(10 * time.Hour)
	}

	return time.Now().Add(duration)
}

func GetRememberMeTokenExpiration() time.Time {
	prop := props.GetString("rememberMeTokenExpiration", "240h")

	//disable token expiration
	if strings.HasPrefix(prop, "0") {
		return time.Time{}
	}

	duration, err := time.ParseDuration(prop)
	if err != nil {
		//Saving the default setting
		log.Println("[WARNING] Invalid 'remember me' token expiration duration. Using default value of 240 hours.")
		props.Set("rememberMeTokenExpiration", "240h")
		Save(props)

		//Return the default value
		return time.Now().Add(10 * time.Hour)
	}

	return time.Now().Add(duration)
}

//GetPrivateKey reads the server's RSA keypair from the file system
func GetPrivateKey() *rsa.PrivateKey {
	// To the fellows developers that are reading this,
	// Check out the following article about saving rsa keypair as pem files:
	// https://medium.com/@Raulgzm/export-import-pem-files-in-go-67614624adc7

	//open the private key file
	pemFile, err := os.Open("data/haddock.pem")
	if err != nil {
		//if the file doesn't exist, generate a new keypair
		if errors.Is(err, os.ErrNotExist) {
			log.Println("[WARNING] Server's private key not found. Generating a new one...")
			return GeneratePrivateKey()
		}
		log.Fatal("Error while opening server's PEM key: ", err)
	}
	defer pemFile.Close()

	//Create a buffer to contain the PEM file's content
	pemFileInfo, _ := pemFile.Stat()
	pemBytes := make([]byte, pemFileInfo.Size())

	buffer := bufio.NewReader(pemFile)
	_, err = buffer.Read(pemBytes)
	if err != nil {
		if GetDebugMode() {
			log.Println("[FATAL] Internal server error while decoding the PEM file: \nError: ", err)
		} else {
			log.Println("[FATAL] Internal server error while decoding the PEM file.")
		}
		os.Exit(1)
	}

	//decode the PEM file
	data, _ := pem.Decode(pemBytes)
	if err != nil {
		if GetDebugMode() {
			log.Println("[FATAL] Internal server error while decoding the server's PEM key: \nError: ", err)
		} else {
			log.Println("[FATAL] Internal server error while decoding the server's PEM key.")
		}
		os.Exit(1)
	}

	//if data is nil then the pem file is invalid, gen a new key
	if data == nil {
		log.Println("[WARNING] Server's private key is invalid or corrupted. Generating a new one...")
		return GeneratePrivateKey()
	}

	//parse the data into a rsa PrivateKey
	privateKey, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		if GetDebugMode() {
			log.Print("[WARNING] Internal server error while parsing the server's PEM key. Generating a new one... \nError: ", err)
		} else {
			log.Print("[WARNING] Internal server error while parsing the server's PEM key. Generating a new one...aaa")
		}

		return GeneratePrivateKey()
	}

	return privateKey
}

//GeneratePrivateKey generates a new RSA keypair and saves it to the file system
func GeneratePrivateKey() *rsa.PrivateKey {
	// To the fellows developers that are reading this,
	// Check out the following article about saving rsa keypair as pem files:
	// https://medium.com/@Raulgzm/export-import-pem-files-in-go-67614624adc7

	//generate a private key using crypto/rsa
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		log.Fatal("Error generating server keypair: ", err)
	}

	// PEM FILE
	// save PEM file (check out more about pem files here: https://en.wikipedia.org/wiki/Privacy-Enhanced_Mail)
	pemFile, err := os.Create("data/haddock.pem")
	if err != nil {
		log.Fatal("Error while saving server's PEM key: ", err)
	}
	defer pemFile.Close()

	// http://golang.org/pkg/encoding/pem/#Block
	var pemKey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}

	if err = pem.Encode(pemFile, pemKey); err != nil {
		log.Fatal("Error while saving server's PEM key: ", err)
	}

	return privateKey
}
