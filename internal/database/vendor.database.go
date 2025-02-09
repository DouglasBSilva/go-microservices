package database

import (
	"context"
	"errors"

	"github.com/DouglasBSilva/go-microservices/internal/dberrors"
	"github.com/DouglasBSilva/go-microservices/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (c Client) GetAllVendors(ctx context.Context) ([]models.Vendor, error) {
	var vendors []models.Vendor

	result := c.DB.WithContext(ctx).Find(&vendors)

	return vendors, result.Error
}

func (c Client) AddVendor(ctx context.Context, vendor *models.Vendor) (*models.Vendor, error) {
	vendor.VendorID = uuid.NewString()

	result := c.DB.WithContext(ctx).Create(&vendor)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}

	return vendor, nil
}

func (c Client) GetVendor(ctx context.Context, vendorId string) (*models.Vendor, error) {
	var vendor models.Vendor
	result := c.DB.WithContext(ctx).
		Where(models.Vendor{VendorID: vendorId}).
		First(&vendor)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{}
		}

		return nil, result.Error
	}

	return &vendor, result.Error
}
