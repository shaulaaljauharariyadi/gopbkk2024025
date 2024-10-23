package main

import "fmt"

func main() {

	bookingName := "booking kita"
	const bookingTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("bookingTickets adalah %T,remainingTickets adalah %T,bookingName adalah %T\n", bookingTickets, remainingTickets, bookingName)

	fmt.Printf("selamat datang di aplikasi %v\n", bookingName)
	fmt.Printf("jumlah tiket tersedia %v tiket dan %v tersisa.\n", bookingTickets, remainingTickets)
	fmt.Println("dapatkan tiket anda disini")

	var firstName string
	var lastName string
	var email string
	var userTicket int
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

	fmt.Printf("user %v %v membeli %v tiket,dan kode konfirmasi akan dikirimkan ke email %v\n", firstName, lastName, userTicket, email)

}
