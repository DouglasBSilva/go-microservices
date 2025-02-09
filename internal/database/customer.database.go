package database

import (
	"context"
	"errors"

	"github.com/DouglasBSilva/go-microservices/internal/dberrors"
	"github.com/DouglasBSilva/go-microservices/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (c Client) GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error) {
	var customers []models.Customer
	result := c.DB.WithContext(ctx).
		Where(models.Customer{Email: emailAddress}).
		Find(&customers)

	return customers, result.Error
}

func (c Client) AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	customer.CustomerID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&customer)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}

		return nil, result.Error
	}

	return customer, nil
}

func (c Client) GetCustomer(ctx context.Context, customerId string) (*models.Customer, error) {
	var customer models.Customer
	result := c.DB.WithContext(ctx).
		Where(models.Customer{CustomerID: customerId}).
		First(&customer)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{}
		}

		return nil, result.Error
	}

	return &customer, result.Error
}
