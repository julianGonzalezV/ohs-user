package request

type UserRequest struct {
	BusinessId  string        `json:"businessId"`
	Sku         string        `json:"sku"`
	Image       string        `json:"image"`
	Name        string        `json:"name"`
	Description string        `json:"description,omitempty"`
	Category    string        `json:"category,omitempty"`
	State       string        `json:"state,omitempty"`
	Price       float64       `json:"price"`
	ProductType string        `json:"type,omitempty"`
	Items       []ItemRequest `json:"items,omitempty"`
}

type ItemRequest struct {
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Category    string `json:"category,omitempty"`
	State       string `json:"state,omitempty"`
}
