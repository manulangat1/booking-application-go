package main

import (
	// "booking-app/helper"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

// import helper;
var conferenceName string = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets uint = 50

type UserData struct {
	name            string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have a total of %v \n", conferenceTickets)
	fmt.Printf("Get your tickets to attend! %v still up for grabs! \n", remainingTickets)

	// arrays in Go have a fixed size and also define the  data type
	// var bookings = [50]string{}
	var bookings []string

	// for len(bookings) < 2 {
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

	validUserInputs := validateUserInput(userName, email, userTickets, remainingTickets)

	if !validUserInputs {
		fmt.Println("Invalid  input")
		// break
	}

	// create a map for a user
	var userData = make(map[string]string)
	// bookings[0] = userName
	userData["userName"] = userName
	userData["email"] = email
	userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	bookings = append(bookings, userName)

	if userTickets > remainingTickets {
		fmt.Println("Invalid input, please select the  correct number of tickets remaining. ")
		// break
		// continue
	}

	remainingTickets = remainingTickets - userTickets

	// userTickets = 2
	fmt.Printf("user %v booked %v tickets %v \n", userName, userTickets, email)

	// remaining tickets
	fmt.Printf("%v remaing tickets for %v \n", remainingTickets, conferenceName)

	fmt.Printf("The whole Slice: %v \n", bookings)
	fmt.Printf("Type of Slice is: %T \n", bookings)
	fmt.Printf("Slice length :v %v \n", len(bookings))

	noTicketsRemaining := remainingTickets <= 0

	firstNames := printFirstNames(bookings, conferenceName)

	fmt.Printf("%v \n", firstNames)
	if noTicketsRemaining {
		//  end the program
		fmt.Println("Conference is booked out. Please come back next year!")
		// break
	}

	wg.Add(1)
	go sendTicket(userTickets, userName)

	wg.Wait()

	// }

	// slice is an abstraction of an array: Create a new array without a size dfn

}

func greetUsers(name string, confName string) {
	fmt.Println("Hi", name, "Welcome to ", confName)
}

func printFirstNames(bookings []string, conferenceName string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstName := names[0]
		fmt.Println("My name is \n", firstName)
		firstNames = append(firstNames, firstName)

		greetUsers(firstName, conferenceName)

	}
	return firstNames
}

func validateUserInput(firstName string, email string, userTickets uint, remainingTickets uint) bool {
	isValidName := len(firstName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidEmail && isValidName && isValidTicketNumber
}

func sendTicket(userTickets uint, userName string) {
	time.Sleep(10 * time.Second)
	// save it to a var named ticket
	var ticket = fmt.Sprintf("%v tickets for %v \n", userTickets, userName)
	fmt.Printf("Sending ticket to email address %v \n", ticket)

	wg.Done()

}
