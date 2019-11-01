package main

import (
	"fmt"
	"os"
)

// returns the user ID for the person who is logged in. Returns -1 if no user is authorized.
// Reads from certificate file and a timestamp to determine if user is logged in or not.
func (a *app) doAuth() int32 {
	user := readUserCredentials()
	if *auth == flagLogin || user == -1 {
		// If user is requesting to login
		user = loginUser()
		if user == -1 {
			// Still invalid credentials
			fmt.Println("Invalid user credentials")
			os.Exit(1)
		}
	}
	return user
}

// readUserCredentials reads the credentials from the cert file, matches them with the database entries
// and returns the UUID of the user.
func readUserCredentials() int32 {
	return 0
}

func loginUser() int32 {
	// First, call logout to logout any current user
	logout()

	// now login the user!
	fmt.Println("Please enter your credentials...")
	fmt.Print("Username: ")
	var username string
	_, err := fmt.Scanln(&username)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	fmt.Println("You entered: " + username)
	return 0
}

func logout() {
	// deletes the certificate file and logs the user out
}
