package iNotificationFactory

import (
	sender "Go-advanced-course/src/OOP/factory/iSender"
)

type INotificationFactory interface {
	SendNotification()
	GetSender() sender.ISender
}
