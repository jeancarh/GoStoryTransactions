package models

type Transaction struct {
	ID          int     `csv:"Id" json:"id" bson:"id"`
	Date        string  `csv:"Date" json:"date" bson:"date"`
	Transaction float64  `csv:"Transaction" json:"transaction" bson:"transaction"`
	TransactionType string `csv:"TransactionType" json:"transactiontype" bson:"transactiontype"`
}