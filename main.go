package main

import (
	"flag"
)

var (
	// Flags
	auth    = flag.String("auth", "", "perform authentication operation (login, logout)")
	list    = flag.String("ls", "", "list entries (a=accounts, b=banks, t=transactions, u=users)")
	new     = flag.String("new", "", "create a new entry (a=account, b=bank, t=transaction, u=user)")
	verbose = flag.Bool("v", false, "verbose mode")
	dRange  = flag.String("f", "", "filter results (e.g. >october, 9/1/2018>3/4/2019)")

	// Other Variables
	// accountFlagOpts     = []string{"a", "account"}
	// bankFlagOpts        = []string{"b", "bank"}
	// transactionFlagOpts = []string{"t", "transaction", "trans"}
	// userFlagOpts        = []string{"u", "user"}
)

const (
	flagLogin       = "login"
	flagLogout      = "logout"
	flagAccount     = "a"
	flagBank        = "b"
	flagTransaction = "t"
	flagUser        = "u"
)

type app struct {
	// App fields
	userID int32
}

func main() {
	// Determine if user is logged in... if not, either login or create an account
	// Determine what the user wishes to accomplish
	flag.Parse()

	// Run the app
	a := app{}
	a.doAuth()
	a.doList()
	a.doNew()
}
