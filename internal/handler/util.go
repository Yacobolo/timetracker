package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Notification represents the structure of the notification data.
type Notification struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

// AddNotificationHeaders adds a custom 'hx-trigger' header to the HTTP response for triggering a notification.
func AddNotificationHeaders(w http.ResponseWriter, notificationContent, notificationType string) error {
	allowedTypes := map[string]bool{
		"info":    true,
		"success": true,
		"error":   true,
	}

	// Validate notification type
	if _, ok := allowedTypes[notificationType]; !ok {
		return fmt.Errorf("invalid notification type: %s. allowed types are info, success, error", notificationType)
	}

	// Create notification data
	notificationData := map[string]Notification{
		"notify": {
			Content: notificationContent,
			Type:    notificationType,
		},
	}

	// Convert notification data to JSON
	notificationJSON, err := json.Marshal(notificationData)
	if err != nil {
		return err
	}

	// Add the custom header
	w.Header().Set("hx-trigger", string(notificationJSON))

	return nil
}
