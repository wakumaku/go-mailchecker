package main

import (
	"fmt"

	mailchecker "github.com/wakumaku/go-mailchecker"
)

func main() {
	if mailchecker.IsValid("an_email_here@emailprovider.com") {
		fmt.Println("Valid!")
		return
	}
	fmt.Println("Not valid...")
}
