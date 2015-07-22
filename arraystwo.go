package main

import (
	"flag"
	"fmt"
	"logging"
)

func Fruits(fruits []string) {

	logging.LogData("We are in the Arrays code for Fruits Function", fruits)
	fmt.Println(fruits)

}

func EcommerceComp(comp []string) {

	logging.LogData("Ecommerce related Array example", comp)
	fmt.Println(comp)

}

func main() {

	flag.Parse()

	data := flag.Arg(0)

	fmt.Println(len(data))
	fruits := []string{"Apple", "Banana", "Guava", "Mango", "Grapes", "Fig"}
	// fruits := data
	logging.LogData("Here is the Fruits Array Data", fruits)
	fmt.Println(fruits)

	comp := []string{"www.flipkart.com", "www.snapdeal.com", "www.ebay.com", "www.bestbuy.com"}
	logging.LogData("Here is the Ecommerece based Data", comp)
	fmt.Println(comp)

	fmt.Println(comp[:2])

}
