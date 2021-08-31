package iNotificationFactory

import (
	sender "Go-advanced-course/src/factory/iSender"
)

type INotificationFactory interface {
	SendNotification()
	GetSender() sender.ISender
}
