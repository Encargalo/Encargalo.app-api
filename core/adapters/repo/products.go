package repo

import (
	"CaliYa/core/domain/models"
	"CaliYa/core/domain/ports"
	"context"
	"log"

	"gorm.io/gorm"
)

type productsRepo struct {
	db *gorm.DB
}

func NewProductsRepo(db *gorm.DB) ports.ProductsRepo {
	return &productsRepo{db}
}

func (p *productsRepo) GetProducts(ctx context.Context) ([]models.ProductsShops, error) {

	var products []models.ProductsShops

	err := p.db.Table("business.shop").Select("*,(select name from products.categories) as categories").Scan(&products).Error
	if err != nil {
		log.Fatal(err)
	}

	return products, nil
}
