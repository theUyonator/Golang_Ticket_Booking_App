package main

import "strings"

//This function validates the inputs entered by the user 
func validateUserInput (firstName string, lastName string, email string, userTickets uint) (bool, bool, bool){
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}