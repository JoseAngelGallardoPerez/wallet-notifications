package service

type Receiver struct {
	// To whom notification will be sent
	To []string

	// How this notification will be sent
	NotificationMethod Method
}

// NewReceiver creates new receiver
func NewReceiver(to []string, method Method) *Receiver {
	return &Receiver{to, method}
}
