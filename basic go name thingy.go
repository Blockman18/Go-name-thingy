package main

import (
	"fmt"
)

func main() {
	var name string
	var last string
	fmt.Print("Whats your name?:")
	fmt.Scan(&name)
	fmt.Print("what is your last name?")
	fmt.Scan(&last)
	fmt.Printf("\nyour name is %s", name)
	fmt.Printf("\nyour last name is %s", last)

}
