package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	conferenceTickets := 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have a total of %v \n", conferenceTickets)
	fmt.Printf("Get your tickets to attend! %v still up for grabs! \n", remainingTickets)

	// arrays in Go have a fixed size and also define the  data type
	// var bookings = [50]string{}
	var bookings []string

	for len(bookings) < 2 {
		var userName string
		var userTickets uint
		var email string

		// ask user for their name
		fmt.Println("Enter your full Name")
		fmt.Scan(&userName)

		fmt.Println("Enter your email")
		fmt.Scan(&email)

		fmt.Println("How many tickets do you want to book?")
		fmt.Scan(&userTickets)

		// bookings[0] = userName
		bookings = append(bookings, userName)

		if userTickets > remainingTickets {
			fmt.Println("Invalid input, please select the  correct number of tickets remaining. ")
			// break
			continue
		}

		remainingTickets = remainingTickets - userTickets

		// userTickets = 2
		fmt.Printf("user %v booked %v tickets %v \n", userName, userTickets, email)

		// remaining tickets
		fmt.Printf("%v remaing tickets for %v \n", remainingTickets, conferenceName)

		fmt.Printf("The whole Slice: %v \n", bookings)
		fmt.Printf("Type of Slice is: %T \n", bookings)
		fmt.Printf("Slice length :v %v \n", len(bookings))

		firstNames := []string{}
		for _, booking := range bookings {
			names := strings.Fields(booking)
			firstName := names[0]
			fmt.Println("My name is \n", firstName)
			firstNames = append(firstNames, firstName)

		}
		fmt.Printf("%v \n", firstNames)

		noTicketsRemaining := remainingTickets <= 0

		if noTicketsRemaining {
			//  end the program
			fmt.Println("Conference is booked out. Please come back next year!")
			break
		}

	}

	// slice is an abstraction of an array: Create a new array without a size dfn

}
