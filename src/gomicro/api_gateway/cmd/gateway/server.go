package main

import "net/http"

func main() {
	// endpoints for frontend to access backend
	// we map url paths to certain functions
	http.HandleFunc("/login", login)
	http.HandleFunc("/customer/payment/authorize", customerPaymentAuth)
	http.HandleFunc("/customer/payment/capture", customerPaymentCapture)
	http.HandleFunc("/customer/ledger", customerLedger)
}

func login(http.ResponseWriter, *http.Request) {

}
func customerPaymentAuth(http.ResponseWriter, *http.Request) {

}
func customerPaymentCapture(http.ResponseWriter, *http.Request) {

}
func customerLedger(http.ResponseWriter, *http.Request) {

}
