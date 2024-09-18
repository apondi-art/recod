package main

import (
	"credit"
	"debit"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const dataFile = "data.txt"

// Save account details (name, account number, balance) to file
func saveAccountDetails(name, accountNumber string, balance int) error {
	data := fmt.Sprintf("%s,%s,%d", name, accountNumber, balance)
	return ioutil.WriteFile(dataFile, []byte(data), 0o644)
}

// Load account details from file
func loadAccountDetails() (string, string, int, error) {
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		// If file doesn't exist, return initial values
		return "", "", 1000, nil
	}

	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return "", "", 0, err
	}

	parts := strings.Split(string(data), ",")
	if len(parts) != 3 {
		return "", "", 0, fmt.Errorf("invalid data format")
	}

	name := parts[0]
	accountNumber := parts[1]
	balance, err := strconv.Atoi(parts[2])
	if err != nil {
		return "", "", 0, err
	}

	return name, accountNumber, balance, nil
}

func main() {
	// Load the existing account details from file
	name, accountNumber, balance, err := loadAccountDetails()
	if err != nil {
		fmt.Println("Error loading account details:", err)
		return
	}

	// If there's no existing data, prompt the user for details
	if name == "" || accountNumber == "" {
		fmt.Println("Enter Your Name:")
		fmt.Scanf("%s", &name)
		fmt.Println("Enter Your Account Number:")
		fmt.Scanf("%s", &accountNumber)
	}

	var action string
	var amountStr string

	fmt.Println("Enter amount:")
	fmt.Scanf("%s", &amountStr)
	fmt.Println("Enter action (debit or credit):")
	fmt.Scanf("%s", &action)

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		fmt.Println("Error: Invalid amount")
		return
	}

	switch action {
	case "credit":
		if amount > balance {
			fmt.Println("The account balance is less than the credit amount")
			return
		}
		balance = credit.Credit(balance, amount)
	case "debit":
		balance = debit.Debit(balance, amount)
	default:
		fmt.Println("Only 'debit' or 'credit' are allowed")
		return
	}

	// Save updated balance and account details to file
	err = saveAccountDetails(name, accountNumber, balance)
	if err != nil {
		fmt.Println("Error saving account details:", err)
		return
	}

	// Display the updated balance
	fmt.Printf("%s, Your Balance for the Account Number %s is %d\n", name, accountNumber, balance)
}
