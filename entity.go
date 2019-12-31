package kanka

// Entity represents the common set of attributes of every Kanka object.
type Entity struct {
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	Entry      string        `json:"entry"`
	Image      string        `json:"image"`
	ImageFull  string        `json:"image_full"`
	ImageThumb string        `json:"image_thumb"`
	IsPrivate  bool          `json:"is_private"`
	EntityID   int           `json:"entity_id"`
	Tags       []interface{} `json:"tags"`
	// CreatedAt  CreatedAt     `json:"created_at"`
	CreatedAt string `json:"created_at"`
	CreatedBy int    `json:"created_by"`
	// UpdatedAt UpdatedAt `json:"updated_at"`
	UpdatedAt string `json:"updated_at"`
	UpdatedBy int    `json:"updated_by"`
}

// CreatedAt details when the entry was created.
type CreatedAt struct {
	Date         string `json:"date"`
	TimezoneType int    `json:"timezone_type"`
	Timezone     string `json:"timezone"`
}

// UpdatedAt details when the entry was last updated.
type UpdatedAt struct {
	Date         string `json:"date"`
	TimezoneType int    `json:"timezone_type"`
	Timezone     string `json:"timezone"`
}
