package main

import (
	"fmt"
	"strings"
)

func main() {

	bookingName := "booking kita"
	const bookingTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("selamat datang di aplikasi %v\n", bookingName)
	fmt.Printf("jumlah tiket tersedia %v tiket dan %v tersisa.\n", bookingTickets, remainingTickets)
	fmt.Println("dapatkan tiket anda disini")

	for {

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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTicket > 0 && userTicket <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTicket

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("user %v %v membeli %v tiket,dan kode konfirmasi akan dikirimkan ke email %v\n", firstName, lastName, userTicket, email)
			fmt.Printf("%v tersisa untuk %v\n", remainingTickets, bookingName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

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

	city := "London"

	switch city {
	case "New York":
		//execute code untuk tiket ke new york
	case "Singapore":
		//execute code untuk tiket ke singapore
	case "London", "berlin":
		//code untuk london & berlin
	case "Hong Kong":
		//code
	default:
		fmt.Print("tidak ada kota yang dipilih")
	}

}
