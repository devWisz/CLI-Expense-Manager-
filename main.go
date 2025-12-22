package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Transaction struct {
	ID          int
	Type        string
	Amount      float64
	Category    string
	Description string
	Date        string
}

var transactions []Transaction
var dataFile = "budget.json"
var nextID = 1

func main() {
	loadData()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("1 Add income")
		fmt.Println("2 Add expense")
		fmt.Println("3 View balance")
		fmt.Println("4 List transactions")
		fmt.Println("5 Edit transaction")
		fmt.Println("6 Delete transaction")
		fmt.Println("7 Monthly summary")
		fmt.Println("8 Exit")
		fmt.Print("Choose option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			addTransaction("income", reader)
		case "2":
			addTransaction("expense", reader)
		case "3":
			showBalance()
		case "4":
			listTransactions()
		case "5":
			editTransaction(reader)
		case "6":
			deleteTransaction(reader)
		case "7":
			monthlySummary()
		case "8":
			saveData()
			fmt.Println("Data saved successfully.")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}

