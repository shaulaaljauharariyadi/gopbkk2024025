package main

import (
	"fmt"
	"gopbkk/helper"
	"sync"
	"time"
)

var bookingName = "booking kita"

const bookingTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	// fungsi untuk user input
	firstName, lastName, email, userTicket := getUserInput()

	// fungsi untuk validate input user
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserinput(firstName, lastName, email, userTicket, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTicket, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTicket, firstName, lastName, email)

		// memanggil fungsi untuk print nama awal
		firstNames := getFirstNames()
		fmt.Printf("nama awal dari pembooking adalah: %v\n", firstNames)

		if remainingTickets == 0 {
			//end program
			fmt.Println("tiket telah habis terjual,datang lain waktu.")
			// break
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
	wg.Wait()

}

func greetUsers() {
	fmt.Printf("selamat datang di aplikasi %v\n", bookingName)
	fmt.Printf("jumlah tiket tersedia %v tiket dan %v tersisa.\n", bookingTickets, remainingTickets)
	fmt.Println("dapatkan tiket anda disini")
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

	// membuat map untuk user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTicket,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List Booking %v\n", bookings)

	fmt.Printf("user %v %v membeli %v tiket,dan kode konfirmasi akan dikirimkan ke email %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tersisa untuk %v\n", remainingTickets, bookingName)
}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(25 * time.Second)
	var ticket = fmt.Sprintf("%v tiket untuk %v %v", userTicket, firstName, lastName)
	fmt.Println("#####################")
	fmt.Printf("mengirim tiket:\n %v \n ke email %v\n", ticket, email)
	fmt.Println("#####################")
	wg.Done()
}
