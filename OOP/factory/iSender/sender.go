package iSender

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}
