# CLI-Expense-Manager-

CLI Expense Manager
Overview

CLI Expense Manager is a simple command-line application written in Go to track your personal finances. It allows you to record incomes and expenses, view your current balance, edit or delete transactions, and see monthly summaries. All data is stored locally in a JSON file (budget.json), so it works completely offline.

This tool is lightweight, standalone, and portable, making it easy to manage your finances directly from your terminal or command prompt.

Features

1. Add income transactions

2. Add expense transactions

3. View total balance

4. List all transactions with details (ID, type, date, amount, category, description)

5. Edit existing transactions

6. Delete transactions

7. Monthly summary of income, expenses, and balance

8. Data stored locally in budget.json

How to Run
Prerequisites

Install Go
 (version 1.23.3 or higher)

Basic knowledge of using terminal or command prompt


Step 1: Clone the repository
git clone https://github.com/<your-username>/<repo-name>.git
cd <repo-name>/src

Make sure main.go is inside the src folder.

Step 2: Build the binary

For Windows:
set CGO_ENABLED=0
set GOOS=windows
set GOARCH=amd64
go build -ldflags="-s -w" -o ../CLI-Expense-Manager.exe


This will create a standalone executable CLI-Expense-Manager.exe in the parent folder.

The binary can run on any Windows machine without installing Go.


Step 3: Run the program
../CLI-Expense-Manager.exe

The program will show a menu with options:

Add income

Add expense

View balance

List transactions

Edit transaction

Delete transaction

Monthly summary

Exit

Enter the number corresponding to the action you want to perform.

Step 4: Data Storage

All transaction data is saved in a local file called budget.json.

The file is automatically updated whenever you add, edit, or delete a transaction.

Make sure the budget.json file stays in the same folder as the executable to preserve your data.

Notes

Only positive numbers are accepted for amounts.

Transaction IDs are automatically assigned and used for editing or deleting transactions.

Leave fields empty when editing to keep existing values unchanged.


License

This project is open-source. You can use, modify, and distribute it freely.
