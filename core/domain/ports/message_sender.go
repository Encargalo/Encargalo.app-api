package ports

type MessageSender interface {
	SendCodeToConfirmPhone(to, name, code string) error
}
