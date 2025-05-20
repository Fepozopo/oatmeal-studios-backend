package service

// AddCustomerLocationInput defines the input for adding a customer location
type AddCustomerLocationInput struct {
	CustomerID int32  `json:"customer_id"`
	Address1   string `json:"address_1"`
	Address2   string `json:"address_2,omitempty"`
	City       string `json:"city"`
	State      string `json:"state"`
	ZipCode    string `json:"zip_code"`
	Phone      string `json:"phone,omitempty"`
	Notes      string `json:"notes,omitempty"`
}
