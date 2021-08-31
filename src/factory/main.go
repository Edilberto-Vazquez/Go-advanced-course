package main

import (
	inf "Go-advanced-course/src/factory/iNotificationFactory"
	"Go-advanced-course/src/factory/methods"
	"fmt"
)

func getNotificationFactory(notificationType string) (inf.INotificationFactory, error) {
	if notificationType == "SMS" {
		return &methods.SMSNotification{}, nil
	}
	if notificationType == "Email" {
		return &methods.EmailNotification{}, nil
	}
	return nil, fmt.Errorf("no notification type")
}

func SendNotification(f inf.INotificationFactory) {
	f.SendNotification()
}

func getMethod(f inf.INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	SendNotification(smsFactory)
	SendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}
