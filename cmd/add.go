package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func makeAddCommand() *cobra.Command {
	loadCmd := &cobra.Command{
		Use:   "add",
		Short: "add a new transaction manually",
		Long:  `TBD`,
		Run:   addTransaction,
	}
	return loadCmd
}

func addTransaction(cmd *cobra.Command, args []string) {
	path, err := getDataPath()
	if err != nil {
		Fatalf("failed to find data path: %v", err)
	}
	now := time.Now()
	fullPath := filepath.Join(path, fmt.Sprintf("%v.csv", now.Year()))
	if _, err = os.Lstat(fullPath); err != nil {
		if !os.IsNotExist(err) {
			Fatalf("could not use %v for file path: %v", fullPath, err)
		}
	}
	fmt.Printf("opening file at: %v\n", fullPath)
	fp, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		Fatalf("failed to open data file: %v", err)
	}
	defer fp.Close()

	toAccount := chooseAccount("Select receiving account")
	fromAccount := chooseAccount("Select sending account")
	amount := getAmount()
	date := getDate()
	reference := getReference()
	category := getCategory()
	description := getDescription()
	writer := csv.NewWriter(fp)
	if err = writer.Write([]string{date, fromAccount, toAccount, amount, reference, category, description}); err != nil {
		Fatalf("failed to write to data file: %v", err)
	}
	writer.Flush()
	pterm.FgLightGreen.Println("added new transaction")
}

func chooseAccount(text string) string {
	options := []string{}
	for _, account := range accounts.Accounts {
		options = append(options, fmt.Sprintf("[%-6s] %s", account.ID, account.Name))
	}
	selectedOption, err := pterm.DefaultInteractiveSelect.
		WithDefaultText(text).
		WithOptions(options).
		Show()
	if err != nil {
		Fatalf("could not choose an account: %v", err)
	}
	for index, item := range options {
		if selectedOption == item {
			return accounts.Accounts[index].ID
		}
	}
	Fatalf("failed to find selected account")
	return ""
}

func getDate() string {
	now := time.Now()
	year := now.Year()
	isCurrentYear, err := pterm.DefaultInteractiveConfirm.Show("Did the transaction take place this year?")
	if err != nil {
		Fatalf("failed to get date: %v", err)
	}
	if !isCurrentYear {
		for {
			yearStr, err := pterm.DefaultInteractiveTextInput.
				WithDefaultText("Year").
				WithMultiLine(false).
				Show()
			if err != nil {
				Fatalf("failed to get year: %v", err)
			}
			year64, err := strconv.ParseInt(yearStr, 10, 64)
			if err != nil {
				pterm.FgLightRed.Printfln("given year is not an integer... try again")
				continue
			}
			year = int(year64)
			if year < 1970 || year > 7000 {
				pterm.FgLightRed.Printfln("given year is invalid... try again")
				continue
			}
			break
		}
	}
	options := []string{}
	month := now.Month()
	for i := 0; i < 12; i++ {
		options = append(options, month.String())
		month = (month + 1) % 13
		if month == 0 {
			month++
		}
	}
	selectedMonth, err := pterm.DefaultInteractiveSelect.
		WithDefaultText("Select month").
		WithOptions(options).
		Show()
	if err != nil {
		Fatalf("could not choose a month: %v", err)
	}

	for {
		dayStr, err := pterm.DefaultInteractiveTextInput.
			WithDefaultText("Day").
			WithMultiLine(false).
			Show()
		if err != nil {
			Fatalf("failed to get day: %v", err)
		}
		day, err := strconv.ParseInt(dayStr, 10, 64)
		if err != nil {
			pterm.FgLightRed.Printfln("given day is not an integer... try again")
			continue
		}
		dateStr := fmt.Sprintf("%v-%v-%v", year, selectedMonth, day)
		date, err := time.Parse("2006-January-02", dateStr)
		if err != nil {
			pterm.FgLightRed.Printfln("failed to parse date [%v]... try again", dateStr)
			continue
		}
		return date.Format(time.DateOnly)
	}
}

func getAmount() string {
	for {
		amountStr, err := pterm.DefaultInteractiveTextInput.
			WithDefaultText("Amount").
			WithMultiLine(false).
			Show()
		if err != nil {
			Fatalf("failed to gather amount: %v", err)
		}
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err == nil {
			return fmt.Sprintf("%.2f", amount)
		}
		pterm.FgLightRed.Printfln("could not parse float value... please try again")
	}
}

func getReference() string {
	referenceID, _ := pterm.DefaultInteractiveTextInput.
		WithDefaultText("Reference ID").
		WithMultiLine(false).
		Show()
	return referenceID
}

func getCategory() string {
	selectedOption, err := pterm.DefaultInteractiveSelect.
		WithDefaultText("Select account").
		WithOptions(config.Categories).
		Show()
	if err != nil {
		Fatalf("could not choose an account: %v", err)
	}
	return selectedOption
}

func getDescription() string {
	description, _ := pterm.DefaultInteractiveTextInput.
		WithDefaultText("Description").
		WithMultiLine(false).
		Show()
	return description
}
