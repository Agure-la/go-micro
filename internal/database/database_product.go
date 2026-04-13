package database

import "go-microservice/internal/models"

func (c Client) GetAllProducts(ctx context.Context, vendorID string) ([]models.Product, error){
	var products []models.Product
	result := c.DB.withContext(ctx).where(models.Product(VendorId: vendorID))
	.Find(&prproducts)
	return products, result.error
}