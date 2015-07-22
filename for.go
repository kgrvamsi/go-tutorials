package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
)

// var log = logrus.New()

func init() {

	/*      O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	        O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	        O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	        O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	        O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	        O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist
	        O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	        O_TRUNC  int = syscall.O_TRUNC  // if possible, truncate file when opened.
	*/

	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY|os.O_SYNC, 0666)

	if err != nil {
		panic(err)
	}

	log.SetLevel(log.DebugLevel)
	log.SetOutput(logFile)

}

func main() {

	i := 10

	for k := 0; k < i; k++ {
		log.Info("From the For Loop Block")
		log.Info("Values started Iterating Ten times as Given")
		fmt.Println("Hey Welcome to Golang ..!!!")
	}
}
