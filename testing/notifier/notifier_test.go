package notifier

import (
	"errors"
	"testing"
)

// A fake implementation for testing
type MockMessageSender struct {
	LastUserID  string
	LastMessage string
	SendError   error
}

func (m *MockMessageSender) Send(userID string, message string) error {
	m.LastUserID = userID
	m.LastMessage = message
	return m.SendError
}

func TestUserNotifier_Welcome(t *testing.T) {
	t.Run("Successful send", func(t *testing.T) {
		mockSender := &MockMessageSender{}
		notifier := &UserNotifier{Sender: mockSender}

		userID := "user-123"
		err := notifier.Welcome(userID)

		if err != nil {
			t.Errorf("expected no error, but got %v", err)
		}

		if mockSender.LastUserID != userID {
			t.Errorf("expected user ID to be %q, but got %q", userID, mockSender.LastUserID)
		}

		expectedMessage := "Welcome to our service!"
		if mockSender.LastMessage != expectedMessage {
			t.Errorf("expected message to be %q, but got %q", expectedMessage, mockSender.LastMessage)
		}
	})

	t.Run("Failed send", func(t *testing.T) {
		expectedErr := errors.New("SMS service down")
		mockSender := &MockMessageSender{SendError: expectedErr}
		notifier := &UserNotifier{Sender: mockSender}

		err := notifier.Welcome("user-456")

		if err == nil {
			t.Error("expected an error, but got nil")
		}

		if err != expectedErr {
			t.Errorf("expected error %v, but got %v", expectedErr, err)
		}
	})
}
