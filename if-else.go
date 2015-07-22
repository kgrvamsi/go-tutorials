package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
	_ "time"
)

func init() {

	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_SYNC|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}

	log.SetLevel(log.DebugLevel)
	log.SetOutput(logFile)
	log.SetFormatter(&log.TextFormatter{})

}

func main() {

	i := "India"

	if i == "India" {
		log.Infof("The user seems to be from India")
		fmt.Println("Vamsi is from India")
	} else {
		log.Infof("We are in the Else blog")
		fmt.Println("Hey he is not from India")
	}

}
