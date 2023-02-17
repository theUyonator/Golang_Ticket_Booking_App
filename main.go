package main

// We import the Format package to utilize it's different functions including Printf, Println
// We import the strings package ro have different functinalities for strings 
import (
	"fmt"
	"strings"
)

// In Go, every functionality of an application must be enclosed in a major function called main 
func main() {
	// For a ticket booking application, we need some key variables and constants including:
	/* 
		Conference Name, The number of Conference Tickets, The number of remaining Tickets
		& and a storage mechanism to hold the bookings 
	*/
	var conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	// The bookings will be held in a slice which is depicted by an [] if this was an array it'll
	// be depicted by [length of array]. The reason a slice is used here is we do not know 
	//how many users are going to book so we can't put a size on the storage. All we know 
	//is that there are 50 tickets but 1 user can buy all 50
	var bookings []string
	
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T and conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")

	// This infinite for loop keeps the application running, prompting you to enter a new booking 
	// each time 
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// ask user for their firstname, lastname, email address and number of tickets 
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets you'd like to purchase: ")
		fmt.Scan(&userTickets)

		//We need to validate the inputs from the user 
		var isValidName = len(firstName) >= 2 && len(lastName) >= 2
		var isValidEmail = strings.Contains(email, "@")
		var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets

		//We need a conditional statement to check if user tickets is greater than remaining tickets 
		//this will fix the edge case of a user ordering > 50 tickets 
		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName + " " + lastName)

			// fmt.Printf("The whole slice: %v\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("Slice type: %T\n", bookings)
			// fmt.Printf("Slice length: %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			//We would like to protect the privacy of the users and only share their first names so this
			//takes an index in the splice and splits the names into first names 
			firstNames := []string{}
			/*To interate over the bookings slice, we use a for loops which gives us access to the 
			indicies and associated elements. Range allows iteration in several data structures including 
			arrays and slices. 
			*/

			for _, booking := range bookings {
				//The package strings is used to perform certain actions on a string. strings.Fields 
				//takes a string and splits the string where there is a space
				var names =	strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			/*
				Since everytime the loop runs we are creating new variable names with the next element
				without explixitly using the index variable, we replace the index with an underscore in the 
				for loop to represent unsued variable. It is called a Blank Identifier 
			*/
			fmt.Printf("The first names of the bookings are: %v\n", firstNames)

			//we want to leave the for loop and end the program when all tickets 
			//are sold out. We do this using an if/else staement paired with the break keyword
			
			if remainingTickets == 0 {
				fmt.Printf("The %v tickets are sold out!. We hope to see you next year. ", conferenceName)
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets \n", remainingTickets, userTickets)
		}
		
	}
}
