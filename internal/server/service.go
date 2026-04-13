package server

type Service struct{
	ServiceID string `gorm:"primaryKey" json:"serviceId"`
	Name      string `json:"name"`
	Price     float32 `gorm:"typenumeric" json:"price`
}