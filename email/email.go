package email

import (
    "gopkg.in/gomail.v2"
    "crypto/tls"
	"os"
    "GoStori/models"
	"strconv"
	"fmt"
)

// Define your email settings
const (
    logoURL= "https://upload.wikimedia.org/wikipedia/commons/e/e3/Stori_logo_vertical.png"   // URL of your logo image 
)


// SendTransactionSummaryEmail sends a summary email.
func SendTransaction(transactions []models.Transaction, summary string, isSummary bool) error {
	emailSender := os.Getenv("emailSender")
    emailPass := os.Getenv("emailPass")
    emailRecipient := os.Getenv("emailRecipient")
    smtpServer := os.Getenv("smtpServer")
    smtpPort := os.Getenv("smtpPort")
	if isSummary == true{
		smtpPortInt, err := strconv.Atoi(smtpPort)
		if err != nil {
			return err
		}
		d := gomail.NewDialer(smtpServer, smtpPortInt, emailSender, emailPass)
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		emailBody := generateSummaryEmailBody(summary,logoURL)
		mailer := gomail.NewMessage()
		mailer.SetHeader("From", emailSender)
		mailer.SetHeader("To", emailRecipient)
		mailer.SetHeader("Subject", "Transaction Summary")
		mailer.SetBody("text/html", emailBody)
		if err := d.DialAndSend(mailer); err != nil {
			return err
		}
	
		return nil
	}else {
		smtpPortInt, err := strconv.Atoi(smtpPort)
		if err != nil {
			return err
		}
		d := gomail.NewDialer(smtpServer, smtpPortInt, emailSender, emailPass)
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		for _, transaction := range transactions {
			emailBody := generateSummaryEmail(transaction, logoURL)
			mailer := gomail.NewMessage()
			mailer.SetHeader("From", emailSender)
			mailer.SetHeader("To", emailRecipient) // Replace with the recipient's email
			mailer.SetHeader("Subject", "Transaction")
			mailer.SetBody("text/html", emailBody)
	
			if err := d.DialAndSend(mailer); err != nil {
				return err
			}
		}
	
		return nil
	}
	
}

func generateSummaryEmail(transaction models.Transaction, logoURL string) string {
	amount := fmt.Sprintf("%.2f", transaction.Transaction)
    return fmt.Sprintf(`
	<html>
	<head>
		<style>
			/* Add any global styles here */
			body {
				font-family: Arial, sans-serif;
				background-color: #f4f4f4;
				padding: 20px;
				text-align: center;
			}
	
			.container {
				background-color: #ffffff;
				border-radius: 5px;
				box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
				max-width: 400px;
				margin: 0 auto;
				padding: 20px;
			}
	
			h2 {
				color: #333;
			}
	
			p {
				color: #777;
			}
	
			img {
				max-width: 100px;
				height: auto;
				display: block;
				margin: 0 auto;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<img src="%s" alt="Your Logo">
			<h2>Transaction Summary</h2>
			<p>Date: %s</p>
			<p>Transaction: %s</p>
			<p>Transaction type: %s</p>
		</div>
	</body>
	</html>
    `, logoURL, transaction.Date, amount, transaction.TransactionType)
}

func generateSummaryEmailBody(summary string, logoURL string) string {
    return fmt.Sprintf(`
	<html>
	<head>
		<style>
			/* Add any global styles here */
			body {
				font-family: Arial, sans-serif;
				background-color: #f4f4f4;
				padding: 20px;
				text-align: center;
			}
	
			.container {
				background-color: #ffffff;
				border-radius: 5px;
				box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
				max-width: 400px;
				margin: 0 auto;
				padding: 20px;
			}
	
			h2 {
				color: #333;
			}
	
			p {
				color: #777;
			}
	
			img {
				max-width: 100px;
				height: auto;
				display: block;
				margin: 0 auto;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<img src="%s" alt="Your Logo">
			<h2>Transaction Summary</h2>
			<p>%s</p>
		</div>
	</body>
	</html>
    `, logoURL, summary)
}