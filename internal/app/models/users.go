package models

// Users represents a user in the system
type Users struct {
	ID        uint   `json:"id,omitempty"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	// Add other relevant user fields here
}
