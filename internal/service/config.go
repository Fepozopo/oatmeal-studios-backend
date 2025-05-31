package service

import (
	"time"

	"github.com/google/uuid"
)

// ------------------ Customer Location Inputs ------------------
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

type UpdateCustomerLocationInput struct {
	ID       int32  `json:"id"`
	Address1 string `json:"address_1"`
	Address2 string `json:"address_2,omitempty"`
	City     string `json:"city"`
	State    string `json:"state"`
	ZipCode  string `json:"zip_code"`
	Phone    string `json:"phone,omitempty"`
	Notes    string `json:"notes,omitempty"`
}

// ------------------ Customer Inputs ------------------
type CreateCustomerInput struct {
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

type UpdateCustomerInput struct {
	ID           int32   `json:"customer_id"`
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

// ------------------ Invoice and Order Inputs ------------------
type CreateInvoiceInput struct {
	InvoiceDate        time.Time `json:"invoice_date"`
	OrderID            int32     `json:"order_id"`
	CustomerID         int32     `json:"customer_id"`
	CustomerLocationID int32     `json:"customer_location_id,omitempty"`
	DueDate            time.Time `json:"due_date"`
	Status             string    `json:"status"`
	Total              float64   `json:"total"`
}

type UpdateInvoiceInput struct {
	ID                 int32     `json:"id"`
	InvoiceDate        time.Time `json:"invoice_date"`
	OrderID            int32     `json:"order_id"`
	CustomerID         int32     `json:"customer_id"`
	CustomerLocationID int32     `json:"customer_location_id,omitempty"`
	DueDate            time.Time `json:"due_date"`
	Status             string    `json:"status"`
	Total              float64   `json:"total"`
}

// ------------------ Order Item Inputs ------------------
type CreateOrderItemInput struct {
	OrderID      int32   `json:"order_id"`
	Sku          string  `json:"sku"`
	Quantity     int32   `json:"quantity"`
	Price        float64 `json:"price"`
	Discount     float64 `json:"discount"`
	ItemTotal    float64 `json:"item_total"`
	PocketNumber int32   `json:"pocket_number"`
}

type UpdateOrderItemInput struct {
	ID           int32   `json:"id"`
	Sku          string  `json:"sku"`
	Quantity     int32   `json:"quantity"`
	Price        float64 `json:"price"`
	Discount     float64 `json:"discount"`
	ItemTotal    float64 `json:"item_total"`
	PocketNumber int32   `json:"pocket_number"`
}

// ------------------ Order Inputs ------------------
type CreateOrderInput struct {
	CustomerID         int32     `json:"customer_id"`
	CustomerLocationID int32     `json:"customer_location_id"`
	OrderDate          time.Time `json:"order_date"`
	Status             string    `json:"status"`
	Type               string    `json:"type"`
	Method             string    `json:"method"`
	ShipDate           time.Time `json:"ship_date"`
	PoNumber           string    `json:"po_number"`
	ShippingCost       float64   `json:"shipping_cost"`
	FreeShipping       bool      `json:"free_shipping"`
	ApplyToCommission  bool      `json:"apply_to_commission"`
	SalesRep           string    `json:"sales_rep"`
	Notes              string    `json:"notes"`
}

type UpdateOrderInput struct {
	ID                 int32     `json:"id"`
	CustomerID         int32     `json:"customer_id"`
	CustomerLocationID int32     `json:"customer_location_id"`
	OrderDate          time.Time `json:"order_date"`
	Status             string    `json:"status"`
	Type               string    `json:"type"`
	Method             string    `json:"method"`
	ShipDate           time.Time `json:"ship_date"`
	PoNumber           string    `json:"po_number"`
	ShippingCost       float64   `json:"shipping_cost"`
	FreeShipping       bool      `json:"free_shipping"`
	ApplyToCommission  bool      `json:"apply_to_commission"`
	SalesRep           string    `json:"sales_rep"`
	Notes              string    `json:"notes"`
}

// ------------------ Planogram Inputs ------------------
type CreatePlanogramInput struct {
	Name       string `json:"name"`
	NumPockets int32  `json:"num_pockets"`
	Notes      string `json:"notes"`
}

type UpdatePlanogramInput struct {
	ID         int32  `json:"id"`
	Name       string `json:"name"`
	NumPockets int32  `json:"num_pockets"`
	Notes      string `json:"notes"`
}

type CreatePlanogramPocketInput struct {
	PlanogramID  int32  `json:"planogram_id"`
	PocketNumber int32  `json:"pocket_number"`
	Category     string `json:"category"`
	ProductID    int32  `json:"product_id"`
}

type UpdatePlanogramPocketInput struct {
	ID        int32  `json:"id"`
	Category  string `json:"category"`
	ProductID int32  `json:"product_id"`
}

type AssignPlanogramToLocationInput struct {
	PlanogramID        int32 `json:"planogram_id"`
	CustomerLocationID int32 `json:"customer_location_id"`
}

type RemovePlanogramFromLocationInput struct {
	PlanogramID        int32 `json:"planogram_id"`
	CustomerLocationID int32 `json:"customer_location_id"`
}

type GetPlanogramPocketByNumberInput struct {
	PlanogramID  int32 `json:"planogram_id"`
	PocketNumber int32 `json:"pocket_number"`
}

// ------------------ Product Inputs ------------------
type CreateProductInput struct {
	Type           string    `json:"type"`
	Sku            string    `json:"sku"`
	Upc            string    `json:"upc"`
	Status         string    `json:"status"`
	Cost           float64   `json:"cost"`
	Price          float64   `json:"price"`
	Envelope       string    `json:"envelope,omitempty"`
	Artist         string    `json:"artist,omitempty"`
	Category       string    `json:"category,omitempty"`
	ReleaseDate    time.Time `json:"release_date,omitempty"`
	LastBoughtDate time.Time `json:"last_bought_date,omitempty"`
	Description    string    `json:"description,omitempty"`
	TextFront      string    `json:"text_front,omitempty"`
	TextInside     string    `json:"text_inside,omitempty"`
}

type UpdateProductInput struct {
	ID             uuid.UUID `json:"id"`
	Type           string    `json:"type"`
	Sku            string    `json:"sku"`
	Upc            string    `json:"upc"`
	Status         string    `json:"status"`
	Cost           float64   `json:"cost"`
	Price          float64   `json:"price"`
	Envelope       string    `json:"envelope,omitempty"`
	Artist         string    `json:"artist,omitempty"`
	Category       string    `json:"category,omitempty"`
	ReleaseDate    time.Time `json:"release_date,omitempty"`
	LastBoughtDate time.Time `json:"last_bought_date,omitempty"`
	Description    string    `json:"description,omitempty"`
	TextFront      string    `json:"text_front,omitempty"`
	TextInside     string    `json:"text_inside,omitempty"`
}

// ------------------ Refresh Token Inputs ------------------
type CreateRefreshTokenInput struct {
	Token  string    `json:"token"`
	UserID uuid.UUID `json:"user_id"`
}

// ------------------ Sales Rep Inputs ------------------
type CreateSalesRepInput struct {
	Status    string `json:"status"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Company   string `json:"company"`
	Address1  string `json:"address1,omitempty"`
	Address2  string `json:"address2,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	ZipCode   string `json:"zip_code,omitempty"`
}

type UpdateSalesRepInput struct {
	ID        int32  `json:"id"`
	Status    string `json:"status"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Company   string `json:"company"`
	Address1  string `json:"address1,omitempty"`
	Address2  string `json:"address2,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	ZipCode   string `json:"zip_code,omitempty"`
}

// ------------------ User Inputs ------------------
type RegisterUserInput struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UpdateUserNameInput struct {
	UserID    uuid.UUID `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

type UpdateUserPasswordInput struct {
	UserID      uuid.UUID `json:"user_id"`
	OldPassword string    `json:"old_password"`
	NewPassword string    `json:"new_password"`
}

type AuthenticateUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ------------------- Inventory Transaction Inputs ------------------
type InsertInventoryTransactionInput struct {
	ProductID uuid.UUID `json:"product_id"`
	Change    int32     `json:"change"`
	Reason    string    `json:"reason"`
	Notes     string    `json:"notes,omitempty"`
}
