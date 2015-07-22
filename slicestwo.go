package main

import (
	"fmt"
	"logging"
)

func Fruits(fruits []string) {
	fmt.Println("These are the List of Fruits:", fruits)
}

func main() {

	logging.LogData("We are in the main func")

	data := make([]string, 5)

	/*data = append(data, "Mango")
	data = append(data, "Apple")
	data = append(data, "Grapes")
	data = append(data, "Orange")
	*/
	data[0] = "Mango"
	data[1] = "Apple"
	data[2] = "Grapes"
	data[3] = "Orange"
	data[4] = "Banana"

	logging.LogData("Here is the Data")
	Fruits(data)

	logging.LogData(data[0])
	logging.LogData("Data is with related to Fruits")
}
