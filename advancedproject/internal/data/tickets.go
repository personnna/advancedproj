package data

import "database/sql"

type Ticket struct {
	ID        int64     `json:"id"`
	UserId    int64     `json:"userId"`
	CreatedAt string    `json:"time"`
	Total     int64     `json:"total"`
	Products  []Product `json:"products"`
}

type Product struct {
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Amount int    `json:"amount"`
}

type TicketModel struct {
	DB *sql.DB
}

func NewTicketModel(db *sql.DB) TicketModel {
	return TicketModel{DB: db}
}
