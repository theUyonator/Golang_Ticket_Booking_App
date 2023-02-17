package main

// We import the Format package to utilize it's different functions including Printf, Println
// We import the strings package ro have different functinalities for strings 
import (
	"fmt"
	"strconv"
)
// For a ticket booking application, we need some key variables and constants including:
/* 
	Conference Name, The number of Conference Tickets, The number of remaining Tickets
	& and a storage mechanism to hold the bookings. These varaibles are defined on the package 
	level to allow all functions access to them
*/
const conferenceTickets uint = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)


// In Go, every functionality of an application must be enclosed in a major function called main 
func main() {

	greetUsers()

	// This infinite for loop keeps the application running, prompting you to enter a new booking 
	// each time 
	for {
		//User input prompt
		firstName, lastName, email, userTickets := getUserInput()

		//We need to validate the inputs from the user 
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		//We need a conditional statement to check if user tickets is greater than remaining tickets 
		//this will fix the edge case of a user ordering > 50 tickets 
		if isValidName && isValidEmail && isValidTicketNumber {
			bookings, remainingTickets = bookTicket(firstName, lastName, email, userTickets)

			//We would like to protect the privacy of the users and only share their first names so this
			//takes the user names inputs and prints out just the first names 
			var firstNames = getFirstNames()
			fmt.Printf("The first names of the bookings are: %v\n", firstNames)

			//we want to leave the for loop and end the program when all tickets 
			//are sold out. We do this using an if/else staement paired with the break keyword
			
			if remainingTickets == 0 {
				fmt.Printf("The %v tickets are sold out!. We hope to see you next year. ", conferenceName)
				break
			}
		} else {
			// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets \n", remainingTickets, userTickets)
			if !isValidName{
				fmt.Println("first name or last name you entered is too short, name must be greater that 2 characters!")
			} 
			if !isValidEmail {
				fmt.Println("email address you entered does not contain an @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
		
	}
}

//This function handles the initial greeting of customers on conference ticket booking site 
func greetUsers () {

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}

//This function takes a user input first name and last name and returns just the first names 
func getFirstNames() []string{
	firstNames := []string{}
	/*To interate over the bookings slice, we use a for loops which gives us access to the 
	indicies and associated elements. Range allows iteration in several data structures including 
	arrays and slices. 
	*/

	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}

	/*
	Since the index variable isn't being used explicitly, we replace the index with an underscore in the 
	for loop to represent unsued variable. It is called a Blank Identifier, this stops the error that will arise 
	otherwise 
	*/
	return firstNames

}

//This function prompts the user to enter inputs
func getUserInput() (string, string, string, uint){
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

	return firstName, lastName, email, userTickets
}

//This function handles the ticket booking 
func bookTicket(firstName string, lastName string, email string, userTickets uint) ([]map[string]string, uint){
	remainingTickets = remainingTickets - userTickets
	
	//create map to store user booking information 

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	//We convert the userTickets into a string bbecause  maps only accept one data type
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)


	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return bookings, remainingTickets
}