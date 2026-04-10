package database

import (
	"context"
	"go-microservice/internal/models"
)

func (c Client) GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error) {
	var customers []models.Customer
	result := c.Db.withContext(ctx).
	where(models.Customer{Email: emailAddress}).
	find(&customers)
	return customers, result.Error
}