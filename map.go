package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
)

func init() {
	logFile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_SYNC|os.O_WRONLY, 0666)

	if err != nil {
		panic(err)
	}

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetLevel(log.DebugLevel)
	log.SetOutput(logFile)

}

func main() {

	i := make(map[string]int)
	log.Info("Here i'm creating a Map variable i")

	i["Ashok"] = 5
	i["Vamsi"] = 6

	log.WithFields(log.Fields{"Ashok": 5, "Vamsi": 6}).Info("Here we have created Array of elements")

	n := map[string]int{"hey": 1}
	fmt.Println(n)
	fmt.Println(i)

}
