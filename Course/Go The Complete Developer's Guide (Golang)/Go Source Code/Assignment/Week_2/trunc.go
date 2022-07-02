package main

import (
	"fmt"
	"log"
)

var price float64

func main() {
	fmt.Println("What is the price of apple?")
	_, err := fmt.Scan(&price)
	if err != nil {
		log.Printf("[Error] Invalid!")
	}

	fmt.Printf("The truncated version of '%v' is '%v.'\n", price, int64(price))
}
