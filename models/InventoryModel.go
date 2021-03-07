package models

type InventoryItem struct {
	ID       int    `json:"ID"`
	Category string `json:"Category"`
	Product  string `json:"Product"`
	Quantity int    `json:"Quantity"`
}

type Inventory struct {
	ID       int    `gorm:"column:ID;primary_key"`
	Category string `gorm:"column:Category"`
	Product  string `gorm:"column:Product"`
	Quantity int    `gorm:"column:Quantity"`
}