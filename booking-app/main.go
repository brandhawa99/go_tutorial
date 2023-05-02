package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type userData struct {
	firstName   string
	lastName    string
	email       string
	ticketCount uint
}

var wg sync.WaitGroup

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings = make([]userData, 0)
	// ^ is a slice not an array slice is  more efficient

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {

		firstName, lastName, email, userTickets := getUserData()
		isValidName, isValidEmail, isValidTicketNumber := validData(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketNumber {

			remainingTickets -= userTickets
			usr := userData{
				firstName:   firstName,
				lastName:    lastName,
				email:       email,
				ticketCount: userTickets,
			}
			bookings = append(bookings, usr)

			fmt.Printf("Thank you  %v %v, for booking %v tickets, you will receive your confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				firstNames = append(firstNames, booking.firstName)

			}
			fmt.Printf("Bookings List: %v\n", firstNames)
			wg.Add(1)
			go sendEmail(usr)
			if remainingTickets == 0 {
				fmt.Println("Our Conference is booked up come back next year")
				break
			}
		} else if userTickets == remainingTickets {
			fmt.Println("You got the last tickets ")
		} else {
			invalidOutPut(isValidEmail, isValidName, isValidTicketNumber)
		}

	}

}

func sendEmail(usr userData) {
	time.Sleep(10 * time.Second)
	fmt.Printf("\nHey, %v %v your %v are send to %v\n", usr.firstName, usr.lastName, usr.ticketCount, usr.email)
	wg.Done()
}

func invalidOutPut(isValidEmail bool, isValidName bool, isValidTicketNumber bool) {
	if isValidEmail {
		fmt.Println("Email data was invalid")
	}
	if isValidName {
		fmt.Println("Name Data was invalid")
	}
	if isValidTicketNumber {
		fmt.Println("Ticket number data is invalid ")
	}
}

func getUserData() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your Last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your Email: ")
	fmt.Scan(&email)

	fmt.Print("Enter How many tickets you want: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func validData(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are left\n", conferenceTickets, remainingTickets)
	fmt.Println("-- Get your tickets here to attend --")
}
