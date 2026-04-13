package database

import "go-microservice/internal/models"

func (c Client) GetAllVendors(ctx context.Context) ([]models.Vendo, error){
	var vendors []models.Vendo
	result := c.Db.withContext(ctx).Find(&vendors)
	return vendors, result.Error
}