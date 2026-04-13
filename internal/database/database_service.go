package database

import "go-microservice/internal/models"

func (c Client) GetAllServices(ctx context.Context) ([]models.Service, error) {
	var services []models.Service
	result := c.Db.withContext(ctx).Find(&services)
	return services, result.error
}