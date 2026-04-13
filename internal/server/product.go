package server

type Product struct{
	ProductID string `gorm:"primaryKey" json:"productId"`
}