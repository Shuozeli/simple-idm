package notification

import (
	"testing"
)

func TestNewNotificationManager(t *testing.T) {
	nm := NewNotificationManager()
	if nm == nil {
		t.Error("NewNotificationManager returned nil")
	}
	if nm.notifiers == nil {
		t.Error("notifiers map not initialized")
	}
	if nm.notificationRegistry == nil {
		t.Error("notificationRegistry map not initialized")
	}
}

func TestRegisterNotifier(t *testing.T) {
	nm := NewNotificationManager()
	mockNotifier := &MockNotifier{}

	// Test registering a notifier
	nm.RegisterNotifier(EmailSystem, mockNotifier)
	if n, exists := nm.notifiers[EmailSystem]; !exists {
		t.Error("Notifier not registered")
	} else if n != mockNotifier {
		t.Error("Wrong notifier registered")
	}

	// Test overwriting existing notifier
	newMockNotifier := &MockNotifier{}
	nm.RegisterNotifier(EmailSystem, newMockNotifier)
	if n := nm.notifiers[EmailSystem]; n != newMockNotifier {
		t.Error("Notifier not overwritten")
	}
}

func TestRegisterNotification(t *testing.T) {
	nm := NewNotificationManager()

	tests := []struct {
		name        string
		notifType   NotificationType
		system      NotificationSystem
		template    string
		shouldError bool
	}{
		{
			name:        "Valid registration",
			notifType:   ExampleNotification,
			system:      EmailSystem,
			template:    "templates/example.tmpl",
			shouldError: false,
		},
		{
			name:        "Empty notification type",
			notifType:   "",
			system:      EmailSystem,
			template:    "templates/example.tmpl",
			shouldError: true,
		},
		{
			name:        "Empty system",
			notifType:   ExampleNotification,
			system:      "",
			template:    "templates/example.tmpl",
			shouldError: true,
		},
		{
			name:        "Empty template",
			notifType:   ExampleNotification,
			system:      EmailSystem,
			template:    "",
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := nm.RegisterNotification(tt.notifType, tt.system, tt.template)
			if tt.shouldError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.shouldError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !tt.shouldError {
				if template, exists := nm.notificationRegistry[tt.notifType][tt.system]; !exists {
					t.Error("Template not registered")
				} else if template != tt.template {
					t.Errorf("Wrong template registered. Got %s, want %s", template, tt.template)
				}
			}
		})
	}
}

func TestSend(t *testing.T) {
	nm := NewNotificationManager()
	mockEmailNotifier := &MockNotifier{}
	mockSMSNotifier := &MockNotifier{}

	// Register notifiers
	nm.RegisterNotifier(EmailSystem, mockEmailNotifier)
	nm.RegisterNotifier(SMSSystem, mockSMSNotifier)

	// Register notifications
	err := nm.RegisterNotification(ExampleNotification, EmailSystem, "templates/example_email.tmpl")
	if err != nil {
		t.Fatalf("Failed to register email notification: %v", err)
	}
	err = nm.RegisterNotification(ExampleNotification, SMSSystem, "templates/example_sms.tmpl")
	if err != nil {
		t.Fatalf("Failed to register SMS notification: %v", err)
	}

	// Test sending notification
	testData := NotificationData{
		To:      "user@example.com",
		Subject: "Test Subject",
		Body:    "Test Body",
	}

	err = nm.Send(ExampleNotification, testData)
	if err != nil {
		t.Errorf("Failed to send notification: %v", err)
	}

	// Verify email notification was sent
	if len(mockEmailNotifier.SentNotifications) != 1 {
		t.Error("Email notification not sent")
	} else {
		sent := mockEmailNotifier.SentNotifications[0]
		if sent.To != testData.To || sent.Subject != testData.Subject || sent.Body != testData.Body {
			t.Error("Email notification data mismatch")
		}
	}

	// Verify SMS notification was sent
	if len(mockSMSNotifier.SentNotifications) != 1 {
		t.Error("SMS notification not sent")
	} else {
		sent := mockSMSNotifier.SentNotifications[0]
		if sent.To != testData.To || sent.Subject != testData.Subject || sent.Body != testData.Body {
			t.Error("SMS notification data mismatch")
		}
	}
}

func TestSendErrors(t *testing.T) {
	nm := NewNotificationManager()

	// Test sending with unregistered notification type
	err := nm.Send("unregistered", NotificationData{})
	if err == nil {
		t.Error("Expected error for unregistered notification type")
	}

	// Register notification without registering notifier
	err = nm.RegisterNotification(ExampleNotification, EmailSystem, "templates/example.tmpl")
	if err != nil {
		t.Fatalf("Failed to register notification: %v", err)
	}

	// Test sending with missing notifier
	err = nm.Send(ExampleNotification, NotificationData{})
	if err == nil {
		t.Error("Expected error for missing notifier")
	} else if err.Error() != "no notifier registered for system: email" {
		t.Errorf("Unexpected error message: %v", err)
	}
}
