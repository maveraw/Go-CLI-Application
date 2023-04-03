package main

import (
	"fmt"
	"sync"
	"time"
)

var confName string = "Go Conference"

const confTickets uint = 50

var remainTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberofTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, userTickets, email, remainTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of our bookings are: %v\n", firstNames)

		if remainTickets == 0 {
			// end program
			fmt.Printf("Our conference is booked out.Come Back next year.")
			//break
		}
	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}

		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")

		}

		if !isValidTicketNumber {
			fmt.Println("number of tickets is invalid. Try again!!")
		}

	}
	wg.Wait()

}

func greetUser() {
	fmt.Printf("Conftickets is of the type %T , confName is of the type %T and remaintikets is of the type %T\n", confTickets, confName, remainTickets)
	fmt.Printf("Welcome to our %v of booking-app\n", confName)
	fmt.Printf("GET YOUR %v tickets available, we only have %v left!!\n", confTickets, remainTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//ask user for their name

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainTickets = remainTickets - userTickets

	// create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberofTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v )for bookig %v tickets. You will receive an email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainTickets, confName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	fmt.Println("###################")
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("Sending ticket:\n %v \n into email address %v\n", ticket, email)
	fmt.Println("###################")
	wg.Done()
}
