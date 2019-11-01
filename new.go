package main

import (
	"fmt"
)

func (a *app) doNew() {
	if *new != "" {
		// User wishes to enter a new item
		switch *new {
		case flagAccount:
			// List accounts associated with this user
		case flagBank:
			// List banks associated with this user
		case flagTransaction:
			// List the last set of transactions for this user
		case flagUser:
			// List the user names available
		default:
			fmt.Println("Sorry... list option not valid")
		}
	}
}

type vUser struct {
	id        string
	firstName string
	lastName  string
	userName  string
	password  string
}

func newUser() {
	// use fmt.Scanln() to gather user details
}
