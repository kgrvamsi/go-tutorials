package main

import (
	"fmt"
	"time"
)

func Multiplication(value int, val int) (result int) {

	res := value * value

	return res

}

func main() {

	messages := make(chan string)

	go func() {

		messages <- "Hello World!!"
	}()

	go func() {

		messages <- "How Are you!!"
	}()

	go func() {

		for i := 0; i <= 400; i++ {
			messages <- "Integer"
		}
	}()

	for {
		time.Sleep(time.Second)
		fmt.Println(<-messages)
	}

	close(messages)

}
