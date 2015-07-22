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

	log.SetOutput(logFile)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main() {

	var i [5]int

	x := 5

	for k := 0; k < x; k++ {
		i[k] = k
		log.Info("Appending the Values to the array")
		fmt.Println("Adding values to the array")
		fmt.Println(i)
	}

	y := [...]string{"Hey", "Hello", "Whazzup"}

	fmt.Println(y)

}
