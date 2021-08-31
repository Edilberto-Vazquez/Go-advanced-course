package methods

import (
	inf "Go-advanced-course/src/factory/iNotificationFactory"
	send "Go-advanced-course/src/factory/iSender"
	"fmt"
)

type EmailNotification struct {
	inf.INotificationFactory
}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending notification via Email")
}

func (EmailNotification) GetSender() send.ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "Gmail"
}
