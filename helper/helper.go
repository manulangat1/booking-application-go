package main

import "strings"

func ValidateUserInput(firstName string, email string, userTickets uint, remainingTickets uint) bool {
	isValidName := len(firstName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidEmail && isValidName && isValidTicketNumber
}
