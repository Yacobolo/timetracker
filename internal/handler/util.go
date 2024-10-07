package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"timetracker/internal/config"
	"timetracker/internal/templates/components"

	"github.com/go-playground/validator/v10"
	"github.com/mattn/go-sqlite3"
)

func Make(h func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("internal server error", "err", err, "path", r.URL.Path)
		}
	}
}

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

func CustomErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Please enter a valid email address"
	default:
		return "Invalid input"
	}
}

type FieldError struct {
	FieldID string
	Error   string
}

func BuildInputFields(fcm config.FieldConfigManager, fieldErrors []FieldError) []components.InputFieldParams {
	// Create a map of field errors for faster lookup
	errorMap := make(map[string]string)
	for _, field := range fieldErrors {
		errorMap[field.FieldID] = field.Error
	}

	// Dynamically create input fields from the field configuration
	var inputFields []components.InputFieldParams
	for _, config := range fcm.Configs {
		inputField := components.InputFieldParams{
			Label:       config.Label,
			Placeholder: config.Placeholder,
			ID:          config.ID,
			Type:        config.Type,
			ErrorMsg:    "", // Default empty error message
		}

		// If an error exists for the field, set the error message
		if errMsg, exists := errorMap[config.ID]; exists {
			inputField.ErrorMsg = errMsg
		}

		inputFields = append(inputFields, inputField)
	}

	return inputFields
}

// Refactored SQLite error handling without fieldConfigs
func handleSQLiteError(err error, fcm *config.FieldConfigManager) []FieldError {
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		if sqliteErr.Code == sqlite3.ErrConstraint {
			if strings.Contains(sqliteErr.Error(), "UNIQUE constraint") {
				// Extract the field name from the error message
				parts := strings.Split(sqliteErr.Error(), ".")
				field_id := parts[len(parts)-1]

				// Check if the field name is in the fieldConfigs
				if field_config, exists := fcm.ConfigMap[field_id]; exists {
					field_label := field_config.Label
					return []FieldError{
						{FieldID: field_id, Error: fmt.Sprintf("That %s already exists", field_label)},
					}
				}
			}
		}
	}
	return nil
}
