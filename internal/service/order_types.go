package service

import (
	"time"
)

type CreateOrUpdateOrderInput struct {
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
