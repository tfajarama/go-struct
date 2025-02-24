package models

type Shipment struct {
	ShipmentID string `json:"shipment_id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Contact    string `json:"contact"`
}
