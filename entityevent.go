package kanka

import "time"

// EntityEvents wraps a list of entity events.
// EntityEvents exists to satisfy the API's JSON structure.
type EntityEvents struct {
	Data []*EntityEvent `json:"data"`
	Sync time.Time      `json:"sync"`
}

// EntityEvent represents a specific calendar event relating to the parent
// entity.
type EntityEvent struct {
	CalendarID     int       `json:"calendar_id"`
	Colour         string    `json:"colour"`
	Comment        string    `json:"comment"`
	CreatedAt      time.Time `json:"created_at"`
	CreatedBy      int       `json:"created_by"`
	Date           string    `json:"date"`
	Day            int       `json:"day"`
	EntityID       int       `json:"entity_id"`
	ID             int       `json:"id"`
	IsPrivate      bool      `json:"is_private"`
	IsRecurring    bool      `json:"is_recurring"`
	Length         int       `json:"length"`
	Month          int       `json:"month"`
	RecurringUntil string    `json:"recurring_until"`
	UpdatedAt      time.Time `json:"updated_at"`
	UpdatedBy      int       `json:"updated_by"`
	Year           int       `json:"year"`
}
