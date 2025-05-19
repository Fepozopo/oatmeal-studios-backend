package service

type CreateOrUpdateCustomerInput struct {
	BusinessName string  `json:"business_name"`
	ContactName  string  `json:"contact_name,omitempty"`
	Email        string  `json:"email,omitempty"`
	Phone        string  `json:"phone,omitempty"`
	Address1     string  `json:"address1,omitempty"`
	Address2     string  `json:"address2,omitempty"`
	City         string  `json:"city,omitempty"`
	State        string  `json:"state,omitempty"`
	ZipCode      string  `json:"zip_code,omitempty"`
	Terms        string  `json:"terms,omitempty"`
	Discount     float64 `json:"discount,omitempty"`
	Commission   float64 `json:"commission,omitempty"`
	SalesRep     string  `json:"sales_rep,omitempty"`
	Notes        string  `json:"notes,omitempty"`
}
