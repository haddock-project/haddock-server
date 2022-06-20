package utils

import (
	"github.com/Kalitsune/Haddock/server/config"
	"log"
)

//HandleError is a function that handles errors differently depending on weather the debugMode is enabled or not.
//returns true if there was an error.
func HandleError(text string, err error) bool {
	//Check if there is an error
	if err != nil {
		//Check if the debugMode is enabled
		if config.GetDebugMode() {
			log.Println(text+"\nDetails:", err)
		} else {
			log.Println(text + ".")
		}
		return true
	}
	return false
}
