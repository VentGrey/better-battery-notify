package notify

import "testing"

func TestSendNotification(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code panicked: %v", r)
		}
	}()
	SendNotification("Test Title", "Test Message", "battery")
}
