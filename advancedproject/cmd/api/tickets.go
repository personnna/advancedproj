package main

import (
	"advancedproject/internal/data"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createTicketHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new ticket")
}

func (app *application) showTicketHandler(w http.ResponseWriter, r *http.Request) {

	ticketId, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	ticket := data.Ticket{
		ID:        ticketId,
		UserId:    123, //i need to change it to the users id
		CreatedAt: time.Now().Format("02 Jan 2006 at 15:04"),
		Total:     104, //write a function to calculate total
		Products: []data.Product{
			data.Product{Name: "cheese", Price: 10, Amount: 2},
			data.Product{Name: "bread", Price: 20, Amount: 1},
			data.Product{Name: "sausage", Price: 30, Amount: 1},
		},
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"ticket": ticket}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
