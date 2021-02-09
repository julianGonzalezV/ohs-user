package entity

// New function is used to create a new struct
func New(items []Item, price float64, businessId, sku, name, description, category, state, productType, image string) *Asset {
	return &Asset{
		BusinessId:  businessId,
		Sku:         sku,
		Image:       image,
		Name:        name,
		Description: description,
		Category:    category,
		State:       state,
		Price:       price,
		ProductType: productType,
		Items:       items,
	}
}

type Asset struct {
	BusinessId  string  `json:"businessId"`
	Sku         string  `json:"sku"`
	Image       string  `json:"image"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Category    string  `json:"category,omitempty"`
	State       string  `json:"state,omitempty"`
	Price       float64 `json:"price"`
	ProductType string  `json:"type,omitempty"`
	Items       []Item  `json:"items,omitempty"`
}

type Item struct {
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Category    string `json:"category,omitempty"`
	State       string `json:"state,omitempty"`
}
