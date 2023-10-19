package main

import (
	"GoStori/csv"
	"GoStori/db"
	"GoStori/email"
	"GoStori/models" // Import the models package
	"fmt"            // Import the fmt package
	"log"
	"time"
	"github.com/joho/godotenv"
)

func main() {
	log.Printf("starting process")
	if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

	// Process the CSV file
	transactions, err := csv.ProcessCSVFile("data.csv")
	if err != nil {
		log.Fatalf("Error processing CSV file: %v", err)
	}

	// Save transactions and accounts to MongoDB
	err = db.SaveToMongoDB(transactions)
	if err != nil {
		log.Fatalf("Error saving to MongoDB: %v", err)
	}

	// Send email summaries for each transaction
	err = email.SendTransaction(transactions, "" , false)
	if err != nil {
		log.Fatalf("Error sending emails: %v", err)
	}

	// Generate and send summary email
	summary := generateSummary(transactions)
	err = email.SendTransaction(nil, summary,true)
	if err != nil {
		log.Fatalf("Error sending summary email: %v", err)
	}

}

func generateSummary(transactions []models.Transaction) string {
	// Initialize variables for summary
	totalBalance := 0.0
	creditTotal := 0.0
	debitTotal := 0.0
	creditCount := 0
	debitCount := 0
	monthTransactionCounts := make(map[string]int)

	// Process transactions
	for _, transaction := range transactions {
		totalBalance += transaction.Transaction

		if transaction.TransactionType == "credit" {
			creditTotal += transaction.Transaction
			creditCount++
		} else {
			debitTotal += transaction.Transaction
			debitCount++
		}
		parsedDate, err := time.Parse("1/2", transaction.Date)
		if err != nil {
			fmt.Printf("Error parsing date: %v\n", err)
			return ""
		}
		parsedMonth := parsedDate.Month()
		month := parsedMonth.String()
		monthTransactionCounts[month]++
	}

	// Generate summary email content
	summary := fmt.Sprintf("Total balance is %.2f\n", totalBalance)
	for month, count := range monthTransactionCounts {
		summary += fmt.Sprintf("Number of transactions in %s: %d\n", month, count)
	}

	if creditCount > 0 {
		averageCreditAmount := creditTotal / float64(creditCount)
		summary += fmt.Sprintf("Average credit amount: %.2f\n", averageCreditAmount)
	}

	if debitCount > 0 {
		averageDebitAmount := debitTotal / float64(debitCount)
		summary += fmt.Sprintf("Average debit amount: %.2f\n", averageDebitAmount)
	}

	return summary
}
