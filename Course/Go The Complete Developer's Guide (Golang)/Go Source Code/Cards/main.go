package main

import "fmt"

func main() {
	cards := deck{"adfas", "wegf"}
	for i, card := range cards {
		fmt.Println(i, card)
	}
}
