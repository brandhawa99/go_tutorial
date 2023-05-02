package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string
	// ^ is a slice not an array slice is  more efficient

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name
		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)
		// get last name
		fmt.Print("Enter your Last name: ")
		fmt.Scan(&lastName)
		// get email
		fmt.Print("Enter your Email: ")
		fmt.Scan(&email)

		fmt.Print("Enter How many tickets you want: ")
		fmt.Scan(&userTickets)

		var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		var isValidEmail bool = strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets

		if isValidEmail && isValidName && isValidTicketNumber {

			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you  %v %v, for booking %v tickets, you will receive your confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])

			}
			fmt.Printf("Bookings List: %v\n", firstNames)
			if remainingTickets == 0 {
				// end program
				fmt.Println("Our Conference is booked up come back next year")
				break
			}
		} else if userTickets == remainingTickets {
			fmt.Println("You got the last tickets ")
		} else {
			if isValidEmail {
				fmt.Println("Email data was invalid")
			}
			if isValidName {
				fmt.Println("Name Data was invalid")
			}
			if isValidTicketNumber {
				fmt.Println("Ticket number data is invalid ")
			}
			fmt.Println("Invalid Input Data")
		}

	}

	city := "London"
	switch city {
	case "Van", "Kelowna":
		fmt.Println("hello from canada")
	case "London":
		fmt.Println("Hello from europe")
	}

}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are left\n", conferenceTickets, remainingTickets)
	fmt.Println("-- Get your tickets here to attend --")
}
