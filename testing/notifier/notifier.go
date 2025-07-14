package notifier

import "fmt"

// The interface our system depends on
type MessageSender interface {
	Send(userID string, message string) error
}

// The component we want to test
type UserNotifier struct {
	Sender MessageSender
}

func (n *UserNotifier) Welcome(userID string) error {
	message := "Welcome to our service!"
	return n.Sender.Send(userID, message)
}

// A real implementation
type SMSSender struct {
	// fields for API keys, etc.
}

func (s *SMSSender) Send(userID string, message string) error {
	// Real logic to call an external SMS API...
	fmt.Printf("Sending SMS to %s: %s\n", userID, message)
	return nil
}
