package main

import (
	"fmt"
	"time"
)

type Notification struct {
	UserId  int
	Massage string
}

func sendEmailAsync(userId int, massage string) {
	// simulate sending an email
	time.Sleep(2 * time.Second)
	fmt.Printf("email notification sent to user %d: %s\n", userId, massage)
}

func main() {
	// simulate multiple user actions that require email notifications
	notifications := []Notification{
		{UserId: 101, Massage: "Your order has been confirmed"},
		{UserId: 202, Massage: "Your account has been created"},
		{UserId: 303, Massage: "Your payment was successful"},
	}

	//process each notification asynchronously using Goroutines
	for _, notification := range notifications {
		go sendEmailAsync(notification.UserId, notification.Massage)
	}

	// main application continues exectuing without waiting for email sending
	fmt.Println("Main application continue...")

	//simulate other operations the main application performs
	time.Sleep(3 * time.Second)

	fmt.Println("Main application finished.")
}
