package methods

import (
	inf "Go-advanced-course/src/factory/iNotificationFactory"
	send "Go-advanced-course/src/factory/iSender"
	"fmt"
)

type SMSNotification struct {
	inf.INotificationFactory
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending notification via SMS")
}

func (SMSNotification) GetSender() send.ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}
