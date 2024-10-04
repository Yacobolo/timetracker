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

// AddHxNotificationTrigger adds or extends a notification in the 'hx-trigger' header of the HTTP response.
func AddHxNotificationTrigger(w http.ResponseWriter, notificationContent, notificationType string) error {
	allowedTypes := map[string]bool{
		"info":    true,
		"success": true,
		"error":   true,
	}

	// Validate notification type
	if _, ok := allowedTypes[notificationType]; !ok {
		return fmt.Errorf("invalid notification type: %s. allowed types are info, success, error", notificationType)
	}

	// Create new notification data
	newNotification := Notification{
		Content: notificationContent,
		Type:    notificationType,
	}

	// Use the helper function to add/extend the 'notify' event in the 'hx-trigger' header
	return AddHxTrigger(w, "notify", newNotification)
}

// AddHxTrigger adds or extends any event (e.g., notification, modal trigger) in the 'hx-trigger' header of the HTTP response.
func AddHxTrigger(w http.ResponseWriter, eventKey string, eventData interface{}) error {
	// Check if hx-trigger header already exists
	existingTrigger := w.Header().Get("hx-trigger")

	// Initialize a map to hold the hx-trigger data
	hxTriggerMap := map[string]interface{}{}

	// If the hx-trigger header already exists, try to unmarshal it into hxTriggerMap
	if existingTrigger != "" {
		if err := json.Unmarshal([]byte(existingTrigger), &hxTriggerMap); err != nil {
			return fmt.Errorf("failed to unmarshal existing hx-trigger: %v", err)
		}
	}

	// Handle empty eventData by assigning an empty map to result in {} in JSON
	if eventData == nil {
		eventData = map[string]interface{}{}
	}

	// Add or update the event in the hxTriggerMap
	hxTriggerMap[eventKey] = eventData

	// Convert the updated map back to JSON
	updatedHxTriggerJSON, err := json.Marshal(hxTriggerMap)
	if err != nil {
		return fmt.Errorf("failed to marshal updated hx-trigger: %v", err)
	}

	// Set the updated hx-trigger header
	w.Header().Set("hx-trigger", string(updatedHxTriggerJSON))

	return nil
}
