package repository

import (
	"context"

	"ApuestaTotal/internal/cart/domain/dto"
	"ApuestaTotal/internal/cart/domain/entity"
	"ApuestaTotal/internal/cart/domain/ports"
	"ApuestaTotal/internal/cart/infrastructure/adapters/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) ports.Cart {
	return &cartRepository{
		db,
	}
}

func (repository cartRepository) GetById(ctx context.Context, id uint) (entity.Cart, error) {
	var modelCart = model.Cart{}
	if result := repository.db.WithContext(ctx).
		Model(&modelCart).
		Preload("Items").
		First(&modelCart, id); result.Error != nil {
		return entity.Cart{}, result.Error
	}

	return modelCart.ToProductDomain(), nil
}

func (repository cartRepository) CreateCart(ctx context.Context, newCart dto.CreateProductDTO) (entity.Cart, error) {
	var modelCart = model.Cart{}
	var modelItem = model.Item{
		Quantity:  newCart.Quantity,
		ProductID: newCart.ProductID,
	}

	if err := repository.db.WithContext(ctx).
		Create(&modelCart).
		Error; err != nil {
		return entity.Cart{}, err
	}

	modelItem.CartID = modelCart.ID

	if err := repository.db.WithContext(ctx).Create(&modelItem).Error; err != nil {
		return entity.Cart{}, err
	}

	if result := repository.db.WithContext(ctx).
		Preload(clause.Associations).
		First(&modelCart, modelCart.ID); result.Error != nil {
		return entity.Cart{}, result.Error
	}

	return modelCart.ToProductDomain(), nil
}

func (repository cartRepository) RemoveCart(ctx context.Context, id uint) error {
	var modelCart = model.Cart{}

	if err := repository.db.WithContext(ctx).
		Select("Items").
		Unscoped().
		Delete(&modelCart, id).Error; err != nil {
		return err
	}
	return nil
}

func (repository cartRepository) AddProduct(ctx context.Context, newProduct dto.AddOrUpdateProductDTO) error {
	var modelItem = model.Item{
		Quantity:  newProduct.Quantity,
		ProductID: newProduct.ProductID,
		CartID:    newProduct.CartID,
	}

	if err := repository.db.WithContext(ctx).Create(&modelItem).Error; err != nil {
		return err
	}

	return nil
}

func (repository cartRepository) RemoveProduct(ctx context.Context, deleteProduct dto.RemoveProductDTO) error {
	var modelItem = model.Item{}

	if err := repository.db.WithContext(ctx).
		Unscoped().
		Where("cart_id = ? AND product_id= ?", deleteProduct.CartID, deleteProduct.ProductID).
		Delete(&modelItem).Error; err != nil {
		return err
	}

	return nil
}

func (repository cartRepository) UpdateProduct(ctx context.Context, updateProduct dto.AddOrUpdateProductDTO) error {
	var modelItem = model.Item{
		ProductID: updateProduct.ProductID,
		CartID:    updateProduct.CartID,
		Quantity:  updateProduct.Quantity,
	}

	if err := repository.db.WithContext(ctx).
		Model(&modelItem).
		Where("cart_id = ? AND product_id = ?", updateProduct.CartID, updateProduct.ProductID).
		Updates(&updateProduct).Error; err != nil {
		return err
	}

	return nil
}
