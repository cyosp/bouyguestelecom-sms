package main

import (
	"flag"
	"fmt"
	"github.com/cyosp/bouyguestelecom-sms/bouyguestelecom"
	"os"
	"time"
)

var version = "0.0.0"

func main() {
	lastname := flag.String("lastname", "", "Last name associated to the Bouygues Telecom account")
	login := flag.String("login", "", "Bouygues Telecom account phone number")
	pass := flag.String("pass", "", "Bouygues Telecom account `password`")
	msg := flag.String("msg", "", "The `message` to send\nEx: \"Hello World\", truncated if more than 160 chars")
	to := flag.String("to", "", "Destination phone numbers\nWith a maximum of 5 numbers, separated by a semicolon and starting by 06 or 07\nEx: \"0601010101;0702020202\"")
	flag.Parse()

	if *login == "" || *pass == "" || *msg == "" || *to == "" {
		fmt.Fprintf(os.Stderr, "Error: 'login', 'password', 'message' and 'to' are required\n\n")
		flag.Usage()
		fmt.Fprintf(os.Stderr, "Using %s %s\n\n", os.Args[0], version)
		os.Exit(1)
	}

	smsLeft, err := bouyguestelecom.SendSms(*lastname, *login, *pass, *msg, *to)
	if err != nil {
		fmt.Printf("Unable to send SMS: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("SMS sent successfully at %s. SMS left: %d.\n", time.Now().Format(time.RFC3339), smsLeft)
}
