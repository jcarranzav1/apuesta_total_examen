package repository

import (
	"context"

	"ApuestaTotal/internal/products/domain/dto"
	"ApuestaTotal/internal/products/domain/entity"
	"ApuestaTotal/internal/products/domain/ports"
	"ApuestaTotal/internal/products/infrastructure/adapters/model"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ports.Product {
	return &productRepository{
		db,
	}
}

func (repository productRepository) GetById(ctx context.Context, id int) (entity.Product, error) {
	var modelProduct model.Product

	if result := repository.db.WithContext(ctx).First(&modelProduct, id); result.Error != nil {
		return entity.Product{}, result.Error
	}

	return modelProduct.ToProductDomain(), nil
}

func (repository productRepository) GetAll(ctx context.Context) ([]entity.Product, error) {
	var modelProducts model.MultipleProduct

	if result := repository.db.WithContext(ctx).Find(&modelProducts); result.Error != nil {
		return []entity.Product{}, result.Error
	}

	return modelProducts.ToProductDomainSlice(), nil
}

func (repository productRepository) Create(ctx context.Context, newProduct dto.ProductCreate) (entity.Product, error) {
	var modelProduct = model.Product{
		Name:  newProduct.Name,
		Price: newProduct.Price,
		Stock: newProduct.Stock,
	}

	if err := repository.db.WithContext(ctx).
		Create(&modelProduct).
		Error; err != nil {

		return entity.Product{}, err
	}
	return modelProduct.ToProductDomain(), nil
}

func (repository productRepository) Update(ctx context.Context, updateProduct dto.ProductUpdate) error {
	var modelProduct model.Product

	if err := repository.db.WithContext(ctx).
		Model(&modelProduct).
		Where("id = ?", updateProduct.ID).
		Updates(&updateProduct).Error; err != nil {
		return err
	}

	return nil
}
