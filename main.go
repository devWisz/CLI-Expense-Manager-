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
			fmt.Println("Your data has been successfully stored.")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}

func addTransaction(tType string, reader *bufio.Reader) {
	amount := readAmount(reader)
	category := readText(reader, "Category: ")
	description := readText(reader, "Description: ")
 
	t := Transaction{
		ID:          nextID,
		Type:        tType,
		Amount:      amount,
		Category:    category,
		Description: description,
		Date:        time.Now().Format("2006-01-02"),
	}

	nextID++
	transactions = append(transactions, t)
	saveData()

	fmt.Println("Transaction added.")
}

func readAmount(reader *bufio.Reader) float64 {
	for {
		fmt.Print("Amount: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		value, err := strconv.ParseFloat(text, 64)
		if err != nil || value <= 0 {
			fmt.Println("Enter a valid positive number.")
			continue
		}
		return value
	}
}

func readText(reader *bufio.Reader, label string) string {
	for {
		fmt.Print(label)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text != "" {
			return text
		}
		fmt.Println("This field cannot be empty.")
	}
}

func showBalance() {
	var income float64
	var expense float64

	for _, t := range transactions {
		if t.Type == "income" {
			income += t.Amount
		} else {
			expense += t.Amount
		}
	}

	fmt.Println()
	fmt.Println("Income :", income)
	fmt.Println("Expense:", expense)
	fmt.Println("Balance:", income-expense)
}

func listTransactions() {
	if len(transactions) == 0 {
		fmt.Println("No transactions found.")
		return
	}

	fmt.Println()
	for _, t := range transactions {
		fmt.Printf(
			"ID %d | %s | %s | %.2f | %s | %s\n",
			t.ID,
			t.Date,
			t.Type,
			t.Amount,
			t.Category,
			t.Description,
		)
	}
}

func editTransaction(reader *bufio.Reader) {
	fmt.Print("Enter transaction ID to edit: ")
	id := readInt(reader)

	for i, t := range transactions {
		if t.ID == id {
			fmt.Println("Leave empty to keep current value.")

			fmt.Print("Enter New amount: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			if text != "" {
				value, err := strconv.ParseFloat(text, 64)
				if err == nil && value > 0 {
					transactions[i].Amount = value
				}
			}

			fmt.Print("New category: ")
			text, _ = reader.ReadString('\n')
			text = strings.TrimSpace(text)
			if text != "" {
				transactions[i].Category = text
			}

			fmt.Print("New description: ")
			text, _ = reader.ReadString('\n')
			text = strings.TrimSpace(text)
			if text != "" {
				transactions[i].Description = text
			}

			saveData()
			fmt.Println("Transaction updated.")
			return
		}
	}

	fmt.Println("Transaction not found.")
}

func deleteTransaction(reader *bufio.Reader) {
	fmt.Print("Enter transaction ID to delete: ")
	id := readInt(reader)

	for i, t := range transactions {
		if t.ID == id {
			transactions = append(transactions[:i], transactions[i+1:]...)
			saveData()
			fmt.Println("Transaction deleted.")
			return
		}
	}

	fmt.Println("Transaction not found.")
}

func readInt(reader *bufio.Reader) int {
	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		value, err := strconv.Atoi(text)
		if err != nil {
			fmt.Print("Enter a valid number: ")
			continue
		}
		return value
	}
}

func monthlySummary() {
	currentMonth := time.Now().Format("2006-01")

	var income float64
	var expense float64

	for _, t := range transactions {
		if strings.HasPrefix(t.Date, currentMonth) {
			if t.Type == "income" {
				income += t.Amount
			} else {
				expense += t.Amount
			}
		}
	}

	fmt.Println()
	fmt.Println("This month income :", income)
	fmt.Println("This month expense:", expense)
	fmt.Println("This month balance:", income-expense)
}

func saveData() {
	data, err := json.MarshalIndent(transactions, "", "  ")
	if err != nil {
		fmt.Println("Failed to save data.")
		return
	}
	_ = os.WriteFile(dataFile, data, 0644)
}

func loadData() {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &transactions)
	if err != nil {
		fmt.Println("Failed to load data.")
		return
	}

	for _, t := range transactions {
		if t.ID >= nextID {
			nextID = t.ID + 1
		}}}     

	
	
