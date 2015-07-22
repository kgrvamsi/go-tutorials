package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
)

func init() {

	logFile, err := os.OpenFile("log.txt", os.O_SYNC|os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		panic(err)
	}

	log.SetOutput(logFile)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main() {

	s := make([]string, 3)

	s[0] = "Hey"
	s[1] = "How"
	s[2] = "Are"

	log.Info("Values added to the Variable s")
	fmt.Println(s)

	log.Info("Adding other Variable in the Slice")
	k := append(s, "You")

	fmt.Println(k)

	// Two Dimensional Slice

	l := make([][]int, 4)

	for x := 0; x < 4; x++ {

		innerLen := x + 1
		l[x] = make([]int, innerLen)
		for y := 0; y < innerLen; y++ {

			l[x][y] = x + y

		}
	}

	fmt.Println(l)

}
