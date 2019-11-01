package main

import (
	"fmt"
)

func (a *app) doList() {
	if *list != "" {
		// User wishes to list out something
		switch *list {
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
