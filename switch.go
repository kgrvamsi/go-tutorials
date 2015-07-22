package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
	"time"
)

func init() {

	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_SYNC|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		panic(err)
	}

	log.SetOutput(logFile)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{})
}

func main() {

	i := time.Now().Unix()

	fmt.Println(i)

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		log.Info("Yupieee Weekend")
		fmt.Println("Weekend Time")
	default:
		log.Info("Hmm Weekday!!")
		fmt.Println("Its a Weekday")
	}

}
