package main

import (
	"fmt"
	"strings"
)

var bookingName = "booking kita"

const bookingTickets int = 50

var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {

		// fungsi untuk user input
		firstName, lastName, email, userTicket := getUserInput()

		// fungsi untuk validate input user
		isValidName, isValidEmail, isValidTicketNumber := validateUserinput(firstName, lastName, email, userTicket)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTicket, firstName, lastName, email)

			// memanggil fungsi untuk print nama awal
			firstNames := getFirstNames()
			fmt.Printf("nama awal dari pembooking adalah: %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("tiket telah habis terjual,datang lain waktu.")
				break
			}
		} else {
			if !isValidName {
				println("nama awal atau akhir yang dimasukkan terlalu pendek!")
			}
			if !isValidEmail {
				println("email yang dimasukkan tidak valid!")
			}
			if !isValidTicketNumber {
				println("jumlah tiket yang dimasukkan tidak valid!")
			}
		}
	}

}

func greetUsers() {
	fmt.Printf("selamat datang di aplikasi %v\n", bookingName)
	fmt.Printf("jumlah tiket tersedia %v tiket dan %v tersisa.\n", bookingTickets, remainingTickets)
	fmt.Println("dapatkan tiket anda disini")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserinput(firstName string, lastName string, email string, userTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint
	//nama depan user
	fmt.Println("Masukkan nama depan: ")
	fmt.Scan(&firstName)
	//nama belakang user
	fmt.Println("Masukkan nama belakang: ")
	fmt.Scan(&lastName)
	//email user
	fmt.Println("Masukkan email: ")
	fmt.Scan(&email)
	//jumlah tiket yang dibeli user
	fmt.Println("ingin memesan berapa tiket: ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTicket(userTicket uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTicket
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("user %v %v membeli %v tiket,dan kode konfirmasi akan dikirimkan ke email %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tersisa untuk %v\n", remainingTickets, bookingName)
}
